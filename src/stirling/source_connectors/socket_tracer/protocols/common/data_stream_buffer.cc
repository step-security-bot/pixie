#include "src/stirling/source_connectors/socket_tracer/protocols/common/data_stream_buffer.h"

#include <algorithm>
#include <deque>
#include <map>
#include <memory>
#include <string>
#include <utility>

#include "src/common/base/base.h"

namespace pl {
namespace stirling {
namespace protocols {

namespace {

// Get element <= key in a map.
template <typename TMapType>
typename TMapType::const_iterator MapLE(const TMapType& map, size_t key) {
  auto iter = map.upper_bound(key);
  if (iter == map.begin()) {
    return map.cend();
  }
  --iter;

  return iter;
}

}  // namespace

void DataStreamBuffer::Reset() {
  buffer_.clear();
  chunks_.clear();
  position_ = 0;
}

// TODO(oazizi): Add checking that the new chunk doesn't overlap with any existing chunk.
//               Return error in such cases.
void DataStreamBuffer::AddNewChunk(size_t pos, size_t size) {
  // Look for the chunks to the left and right of this new chunk.
  auto r_iter = chunks_.lower_bound(pos);
  auto l_iter = r_iter;
  if (l_iter != chunks_.begin()) {
    --l_iter;
  }

  // Does this chunk fuse with the chunk on the left of it?
  bool left_fuse = false;
  if (l_iter != chunks_.end()) {
    size_t l_pos = l_iter->first;
    size_t l_size = l_iter->second;

    left_fuse = (l_pos + l_size == pos);
  }

  // Does this chunk fuse with the chunk on the right of it?
  bool right_fuse = false;
  if (r_iter != chunks_.end()) {
    size_t r_pos = r_iter->first;

    right_fuse = (pos + size == r_pos);
  }

  if (left_fuse && right_fuse) {
    // The new chunk bridges two previously separate chunks together.
    // Keep the left one and increase its size to cover all three chunks.
    l_iter->second += (size + r_iter->second);
    chunks_.erase(r_iter);
  } else if (left_fuse) {
    // Merge new chunk directly to the one on its left.
    l_iter->second += size;
  } else if (right_fuse) {
    // Merge new chunk into the one on its right.
    // Since its key changes, this requires removing and re-inserting the node.
    auto node = chunks_.extract(r_iter);
    node.key() = pos;
    node.mapped() += size;
    chunks_.insert(std::move(node));
  } else {
    // No fusing, so just add the new chunk.
    chunks_[pos] = size;
  }
}

void DataStreamBuffer::AddNewTimestamp(size_t pos, uint64_t timestamp) {
  timestamps_[pos] = timestamp;
}

void DataStreamBuffer::Add(size_t pos, std::string_view data, uint64_t timestamp) {
  if (data.size() > capacity_) {
    size_t oversize_amount = data.size() - capacity_;
    data.remove_prefix(oversize_amount);
    pos += oversize_amount;
  }

  // Calculate physical positions (ppos) where the data would live in the physical buffer.
  ssize_t ppos_front = pos - position_;
  ssize_t ppos_back = pos + data.size() - position_;

  if (ppos_back < 0) {
    // Case 1: Data being added is too far back. Just ignore it.

    // This has been observed to happen a lot on initial deployment,
    // where a large batch of events, with cumulative size greater than the buffer size
    // arrive in scrambled order.
    VLOG(1) << absl::Substitute(
        "Ignoring event that has already been skipped [event pos=$0, current pos=$1].", pos,
        position_);
    return;
  } else if (ppos_front < 0) {
    // Case 2: Data being added is straddling the front-side of the buffer. Cut-off the prefix.

    ssize_t prefix = 0 - ppos_front;
    data.remove_prefix(prefix);
    pos += prefix;
    ppos_front = 0;
  } else if (ppos_back > static_cast<ssize_t>(buffer_.size())) {
    // Case 3: Data being added extends the buffer. Resize the buffer.

    if (pos > position_ + kDataStreamBufferCapacity) {
      // This has been observed to happen a lot on initial deployment,
      // where a large batch of events, with cumulative size greater than the buffer size
      // arrive in scrambled order.
      VLOG(1) << absl::Substitute("Event skips ahead *a lot* [event pos=$0, current pos=$1].", pos, position_);
    }

    ssize_t logical_size = pos + data.size() - position_;
    if (logical_size > static_cast<ssize_t>(capacity_)) {
      // The movement of the buffer position will cause some bytes to "fall off",
      // remove those now.
      size_t remove_count = logical_size - capacity_;
      RemovePrefix(remove_count);
      ppos_front -= remove_count;
      ppos_back -= remove_count;
    }

    DCHECK_GE(ppos_front, 0);
    DCHECK_LE(ppos_front, capacity_);

    DCHECK_GE(ppos_back, 0);
    DCHECK_LE(ppos_back, capacity_);

    ssize_t extension = ppos_back - buffer_.size();
    DCHECK_GE(extension, 0);
    DCHECK_LE(extension, capacity_);

    buffer_.resize(buffer_.size() + extension);
    DCHECK_GE(buffer_.size(), 0);
    DCHECK_LE(buffer_.size(), capacity_);
  } else {
    // Case 4: Data being added is completely within the buffer. Write it directly.

    // No adjustments required.
  }

  // Now copy the data into the buffer.
  memcpy(buffer_.data() + ppos_front, data.data(), data.size());

  // Update the metadata.
  AddNewChunk(pos, data.size());
  AddNewTimestamp(pos, timestamp);
}

std::map<size_t, size_t>::const_iterator DataStreamBuffer::GetChunkForPos(size_t pos) const {
  // Get chunk which is <= pos.
  auto iter = MapLE(chunks_, pos);
  if (iter == chunks_.cend()) {
    return chunks_.cend();
  }

  DCHECK_GE(pos, iter->first);

  // Does the chunk include pos? If not, return {}.
  ssize_t available = iter->second - (pos - iter->first);
  if (available <= 0) {
    return chunks_.cend();
  }

  return iter;
}

std::string_view DataStreamBuffer::Get(size_t pos) const {
  auto iter = GetChunkForPos(pos);
  if (iter == chunks_.cend()) {
    return {};
  }

  size_t chunk_pos = iter->first;
  size_t chunk_size = iter->second;

  ssize_t bytes_available = chunk_size - (pos - chunk_pos);
  DCHECK_GT(bytes_available, 0);

  DCHECK_GE(pos, position_);
  size_t ppos = pos - position_;
  DCHECK_LT(ppos, buffer_.size());
  return std::string_view(buffer_.data() + ppos, bytes_available);
}

StatusOr<uint64_t> DataStreamBuffer::GetTimestamp(size_t pos) const {
  // Ensure the specified time corresponds to a real chunk.
  if (GetChunkForPos(pos) == chunks_.cend()) {
    return error::Internal("Specified position not found");
  }

  // Get chunk which is <= pos.
  auto iter = MapLE(timestamps_, pos);
  if (iter == timestamps_.cend()) {
    LOG(DFATAL) << absl::Substitute(
        "Specified position should have been found, since we verified we are not in a chunk gap "
        "[position=$0]\n$1.",
        pos, DebugInfo());
    return error::Internal("Specified position not found.");
  }

  DCHECK_GE(pos, iter->first);

  return iter->second;
}

void DataStreamBuffer::CleanupMetadata() {
  CleanupChunks();
  CleanupTimestamps();
}

void DataStreamBuffer::CleanupChunks() {
  // Find and remove irrelevant metadata in `chunks_`.

  // Get chunk which is <= position_.
  auto iter = MapLE(chunks_, position_);
  if (iter == chunks_.cend()) {
    return;
  }

  size_t chunk_pos = iter->first;
  size_t chunk_size = iter->second;

  DCHECK_GE(position_, chunk_pos);
  ssize_t available = chunk_size - (position_ - chunk_pos);

  if (available <= 0) {
    // position_ was in a gap area between two chunks, so go back to the next chunk.
    ++iter;
    chunks_.erase(chunks_.begin(), iter);
  } else {
    // Remove all chunks entirely before position_.
    chunks_.erase(chunks_.begin(), iter);

    // Adjust the first chunk's size.
    DCHECK(!chunks_.empty());
    auto node = chunks_.extract(chunks_.begin());
    node.key() = position_;
    node.mapped() = available;
    chunks_.insert(std::move(node));
  }
}

void DataStreamBuffer::CleanupTimestamps() {
  // Find and remove irrelevant metadata in `timestamps_`.

  // Get timestamp which is <= position_.
  auto iter = MapLE(timestamps_, position_);
  if (iter == timestamps_.cend()) {
    return;
  }

  // We are now at the timestamp that covers position_,
  // anything before this is expired and can be removed.
  timestamps_.erase(timestamps_.begin(), iter);

  DCHECK(!timestamps_.empty());
}

void DataStreamBuffer::RemovePrefix(ssize_t n) {
  // Check for positive values of n.
  // For safety in production code, just return.
  DCHECK_GE(n, 0);
  if (n < 0) {
    return;
  }

  buffer_.erase(0, n);
  position_ += n;

  CleanupMetadata();
}

void DataStreamBuffer::Trim() {
  if (chunks_.empty()) {
    return;
  }

  auto& chunk_pos = chunks_.begin()->first;
  DCHECK_GE(chunk_pos, position_);
  size_t trim_size = chunk_pos - position_;

  buffer_.erase(0, trim_size);
  position_ += trim_size;
}

std::string DataStreamBuffer::DebugInfo() const {
  std::string s;

  absl::StrAppend(&s, absl::Substitute("Position: $0\n", position_));
  absl::StrAppend(&s, absl::Substitute("BufferSize: $0/$1\n", buffer_.size(), capacity_));
  absl::StrAppend(&s, "Chunks:\n");
  for (const auto& [pos, size] : chunks_) {
    absl::StrAppend(&s, absl::Substitute("  position:$0 size:$1\n", pos, size));
  }
  absl::StrAppend(&s, "Timestamps:\n");
  for (const auto& [pos, timestamp] : timestamps_) {
    absl::StrAppend(&s, absl::Substitute("  position:$0 timestamp:$1\n", pos, timestamp));
  }
  absl::StrAppend(&s, absl::Substitute("Buffer: $0\n", buffer_));

  return s;
}

}  // namespace protocols
}  // namespace stirling
}  // namespace pl
