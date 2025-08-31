package base58

import (
	"bytes"
	"fmt"
	"io"
	"testing"

	"github.com/dromara/dongle/mock"
)

// BenchmarkStdEncoder_Encode benchmarks the standard base58 encoder with small data
func BenchmarkStdEncoder_Encode(b *testing.B) {
	data := []byte("Hello, World! This is a test string for base58 encoding benchmark.")
	encoder := NewStdEncoder()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		encoder.Encode(data)
	}
}

// BenchmarkStdEncoder_EncodeLarge benchmarks the standard base58 encoder with large data
func BenchmarkStdEncoder_EncodeLarge(b *testing.B) {
	// Create a larger data set for testing
	data := bytes.Repeat([]byte("Hello, World! "), 1000) // ~15KB
	encoder := NewStdEncoder()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		encoder.Encode(data)
	}
}

// BenchmarkStdEncoder_EncodeBinary benchmarks the standard base58 encoder with binary data
func BenchmarkStdEncoder_EncodeBinary(b *testing.B) {
	// Create binary data for testing
	data := make([]byte, 1024)
	for i := range data {
		data[i] = byte(i % 256)
	}
	encoder := NewStdEncoder()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		encoder.Encode(data)
	}
}

// BenchmarkStdEncoder_EncodeEmpty benchmarks the standard base58 encoder with empty data
func BenchmarkStdEncoder_EncodeEmpty(b *testing.B) {
	data := []byte{}
	encoder := NewStdEncoder()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		encoder.Encode(data)
	}
}

// BenchmarkStdEncoder_EncodeSingleByte benchmarks the standard base58 encoder with single byte
func BenchmarkStdEncoder_EncodeSingleByte(b *testing.B) {
	data := []byte{0x41} // 'A'
	encoder := NewStdEncoder()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		encoder.Encode(data)
	}
}

// BenchmarkStdEncoder_EncodeLeadingZeros benchmarks the standard base58 encoder with leading zeros
func BenchmarkStdEncoder_EncodeLeadingZeros(b *testing.B) {
	data := []byte{0x00, 0x00, 0x00, 0x01, 0x02, 0x03, 0x04, 0x05}
	encoder := NewStdEncoder()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		encoder.Encode(data)
	}
}

// BenchmarkStdEncoder_EncodeAllZeros benchmarks the standard base58 encoder with all zeros
func BenchmarkStdEncoder_EncodeAllZeros(b *testing.B) {
	data := []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	encoder := NewStdEncoder()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		encoder.Encode(data)
	}
}

// BenchmarkStdEncoder_EncodeLargeNumber benchmarks the standard base58 encoder with large number
func BenchmarkStdEncoder_EncodeLargeNumber(b *testing.B) {
	data := []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}
	encoder := NewStdEncoder()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		encoder.Encode(data)
	}
}

// BenchmarkStdEncoder_EncodeUnicode benchmarks the standard base58 encoder with unicode data
func BenchmarkStdEncoder_EncodeUnicode(b *testing.B) {
	data := []byte("你好世界，这是一个包含中文的测试字符串")
	encoder := NewStdEncoder()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		encoder.Encode(data)
	}
}

// BenchmarkStdEncoder_EncodeMixedData benchmarks the standard base58 encoder with mixed data types
func BenchmarkStdEncoder_EncodeMixedData(b *testing.B) {
	data := []byte{
		0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07,
		0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E, 0x0F,
	}
	encoder := NewStdEncoder()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		encoder.Encode(data)
	}
}

// BenchmarkStdEncoder_EncodeShortStrings benchmarks the standard base58 encoder with short strings
func BenchmarkStdEncoder_EncodeShortStrings(b *testing.B) {
	shortStrings := [][]byte{
		[]byte("A"),
		[]byte("AB"),
		[]byte("ABC"),
		[]byte("ABCD"),
		[]byte("ABCDE"),
	}
	encoder := NewStdEncoder()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, str := range shortStrings {
			encoder.Encode(str)
		}
	}
}

// BenchmarkStdEncoder_EncodeBlockSize benchmarks the standard base58 encoder with block size data
func BenchmarkStdEncoder_EncodeBlockSize(b *testing.B) {
	// Test with different block sizes to see performance characteristics
	blockSizes := []int{1, 2, 3, 4, 5, 6, 7, 8, 16, 32, 64, 128, 256, 512, 1024}
	encoder := NewStdEncoder()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, size := range blockSizes {
			data := make([]byte, size)
			for j := range data {
				data[j] = byte(j % 256)
			}
			encoder.Encode(data)
		}
	}
}

// BenchmarkStdEncoder_EncodeRandomData benchmarks the standard base58 encoder with random-like data
func BenchmarkStdEncoder_EncodeRandomData(b *testing.B) {
	// Create data that simulates random binary data
	data := make([]byte, 1024)
	for i := range data {
		// Use a simple pattern that creates varied byte values
		data[i] = byte((i*7 + 13) % 256)
	}
	encoder := NewStdEncoder()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		encoder.Encode(data)
	}
}

// BenchmarkStdEncoder_EncodeRepeatedPattern benchmarks the standard base58 encoder with repeated patterns
func BenchmarkStdEncoder_EncodeRepeatedPattern(b *testing.B) {
	// Create data with repeated patterns
	pattern := []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08}
	data := bytes.Repeat(pattern, 128) // 1KB of repeated pattern
	encoder := NewStdEncoder()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		encoder.Encode(data)
	}
}

// BenchmarkStdDecoder_Decode benchmarks the standard base58 decoder with small data
func BenchmarkStdDecoder_Decode(b *testing.B) {
	// Create encoded data for testing
	encoder := NewStdEncoder()
	original := []byte("Hello, World! This is a test string for base58 decoding benchmark.")
	encoded := encoder.Encode(original)
	decoder := NewStdDecoder()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		decoder.Decode(encoded)
	}
}

// BenchmarkStdDecoder_DecodeLarge benchmarks the standard base58 decoder with large data
func BenchmarkStdDecoder_DecodeLarge(b *testing.B) {
	// Create large encoded data for testing
	encoder := NewStdEncoder()
	original := bytes.Repeat([]byte("Hello, World! "), 1000) // ~15KB
	encoded := encoder.Encode(original)
	decoder := NewStdDecoder()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		decoder.Decode(encoded)
	}
}

// BenchmarkStdDecoder_DecodeBinary benchmarks the standard base58 decoder with binary data
func BenchmarkStdDecoder_DecodeBinary(b *testing.B) {
	// Create binary encoded data for testing
	encoder := NewStdEncoder()
	original := make([]byte, 1024)
	for i := range original {
		original[i] = byte(i % 256)
	}
	encoded := encoder.Encode(original)
	decoder := NewStdDecoder()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		decoder.Decode(encoded)
	}
}

// BenchmarkStdDecoder_DecodeEmpty benchmarks the standard base58 decoder with empty data
func BenchmarkStdDecoder_DecodeEmpty(b *testing.B) {
	data := []byte{}
	decoder := NewStdDecoder()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		decoder.Decode(data)
	}
}

// BenchmarkStdDecoder_DecodeLeadingOnes benchmarks the standard base58 decoder with leading ones
func BenchmarkStdDecoder_DecodeLeadingOnes(b *testing.B) {
	// Create encoded data with leading ones (representing leading zeros)
	encoder := NewStdEncoder()
	original := []byte{0x00, 0x00, 0x00, 0x01, 0x02, 0x03}
	encoded := encoder.Encode(original)
	decoder := NewStdDecoder()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		decoder.Decode(encoded)
	}
}

// BenchmarkStdDecoder_DecodeAllOnes benchmarks the standard base58 decoder with all ones
func BenchmarkStdDecoder_DecodeAllOnes(b *testing.B) {
	// Create encoded data with all ones (representing all zeros)
	encoder := NewStdEncoder()
	original := []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	encoded := encoder.Encode(original)
	decoder := NewStdDecoder()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		decoder.Decode(encoded)
	}
}

// BenchmarkStdDecoder_DecodeUnicode benchmarks the standard base58 decoder with unicode data
func BenchmarkStdDecoder_DecodeUnicode(b *testing.B) {
	// Create encoded unicode data for testing
	encoder := NewStdEncoder()
	original := []byte("你好世界，这是一个包含中文的测试字符串")
	encoded := encoder.Encode(original)
	decoder := NewStdDecoder()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		decoder.Decode(encoded)
	}
}

// BenchmarkStdDecoder_DecodeShortStrings benchmarks the standard base58 decoder with short strings
func BenchmarkStdDecoder_DecodeShortStrings(b *testing.B) {
	// Create encoded short strings for testing
	encoder := NewStdEncoder()
	shortStrings := [][]byte{
		[]byte("A"),
		[]byte("AB"),
		[]byte("ABC"),
		[]byte("ABCD"),
		[]byte("ABCDE"),
	}

	var encodedStrings [][]byte
	for _, str := range shortStrings {
		encodedStrings = append(encodedStrings, encoder.Encode(str))
	}

	decoder := NewStdDecoder()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, encoded := range encodedStrings {
			decoder.Decode(encoded)
		}
	}
}

// BenchmarkStdDecoder_DecodeBlockSize benchmarks the standard base58 decoder with block size data
func BenchmarkStdDecoder_DecodeBlockSize(b *testing.B) {
	// Test with different block sizes to see performance characteristics
	blockSizes := []int{1, 2, 3, 4, 5, 6, 7, 8, 16, 32, 64, 128, 256, 512, 1024}
	encoder := NewStdEncoder()
	decoder := NewStdDecoder()

	// Pre-encode all test data
	var encodedData [][]byte
	for _, size := range blockSizes {
		data := make([]byte, size)
		for j := range data {
			data[j] = byte(j % 256)
		}
		encodedData = append(encodedData, encoder.Encode(data))
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, encoded := range encodedData {
			decoder.Decode(encoded)
		}
	}
}

// BenchmarkStdDecoder_DecodeRandomData benchmarks the standard base58 decoder with random-like data
func BenchmarkStdDecoder_DecodeRandomData(b *testing.B) {
	// Create encoded random-like data for testing
	encoder := NewStdEncoder()
	data := make([]byte, 1024)
	for i := range data {
		data[i] = byte((i*7 + 13) % 256)
	}
	encoded := encoder.Encode(data)
	decoder := NewStdDecoder()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		decoder.Decode(encoded)
	}
}

// BenchmarkStdDecoder_DecodeRepeatedPattern benchmarks the standard base58 decoder with repeated patterns
func BenchmarkStdDecoder_DecodeRepeatedPattern(b *testing.B) {
	// Create encoded repeated pattern data for testing
	encoder := NewStdEncoder()
	pattern := []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08}
	data := bytes.Repeat(pattern, 128) // 1KB of repeated pattern
	encoded := encoder.Encode(data)
	decoder := NewStdDecoder()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		decoder.Decode(encoded)
	}
}

// BenchmarkStreamEncoder_Write benchmarks the streaming base58 encoder
func BenchmarkStreamEncoder_Write(b *testing.B) {
	data := []byte("Hello, World! This is a test string for streaming base58 encoding benchmark.")
	var buf bytes.Buffer
	encoder := NewStreamEncoder(&buf)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf.Reset()
		encoder.Write(data)
		encoder.Close()
	}
}

// BenchmarkStreamEncoder_WriteLarge benchmarks the streaming base58 encoder with large data
func BenchmarkStreamEncoder_WriteLarge(b *testing.B) {
	data := bytes.Repeat([]byte("Hello, World! "), 1000) // ~15KB
	var buf bytes.Buffer
	encoder := NewStreamEncoder(&buf)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf.Reset()
		encoder.Write(data)
		encoder.Close()
	}
}

// BenchmarkStreamEncoder_WriteChunked benchmarks the streaming base58 encoder with chunked writes
func BenchmarkStreamEncoder_WriteChunked(b *testing.B) {
	chunks := [][]byte{
		[]byte("Hello, "),
		[]byte("World! "),
		[]byte("This is a "),
		[]byte("test string "),
		[]byte("for streaming "),
		[]byte("base58 encoding "),
		[]byte("benchmark."),
	}
	var buf bytes.Buffer
	encoder := NewStreamEncoder(&buf)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf.Reset()
		for _, chunk := range chunks {
			encoder.Write(chunk)
		}
		encoder.Close()
	}
}

// BenchmarkStreamDecoder_Read benchmarks the streaming base58 decoder
func BenchmarkStreamDecoder_Read(b *testing.B) {
	// Create encoded data for testing
	encoder := NewStdEncoder()
	original := []byte("Hello, World! This is a test string for streaming base58 decoding benchmark.")
	encoded := encoder.Encode(original)

	// Create a reader from the encoded data
	reader := mock.NewFile(encoded, "test.bin")
	decoder := NewStreamDecoder(reader)

	// Buffer to read into
	buf := make([]byte, 1024)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reader.Seek(0, 0) // Reset reader position
		decoder.Read(buf)
	}
}

// BenchmarkStreamDecoder_ReadLarge benchmarks the streaming base58 decoder with large data
func BenchmarkStreamDecoder_ReadLarge(b *testing.B) {
	// Create large encoded data for testing
	encoder := NewStdEncoder()
	original := bytes.Repeat([]byte("Hello, World! "), 1000) // ~15KB
	encoded := encoder.Encode(original)

	// Create a reader from the encoded data
	reader := mock.NewFile(encoded, "test.bin")
	decoder := NewStreamDecoder(reader)

	// Buffer to read into
	buf := make([]byte, 1024)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reader.Seek(0, 0) // Reset reader position
		decoder.Read(buf)
	}
}

// BenchmarkConvenience_Encode benchmarks the convenience Encode function
func BenchmarkConvenience_Encode(b *testing.B) {
	data := []byte("Hello, World! This is a test string for base58 convenience encoding benchmark.")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Encode(data)
	}
}

// BenchmarkConvenience_Decode benchmarks the convenience Decode function
func BenchmarkConvenience_Decode(b *testing.B) {
	// Create encoded data for testing
	encoder := NewStdEncoder()
	original := []byte("Hello, World! This is a test string for base58 convenience decoding benchmark.")
	encoded := encoder.Encode(original)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Decode(encoded)
	}
}

// BenchmarkStdEncoder_EncodeWithError benchmarks the standard base58 encoder with existing error
func BenchmarkStdEncoder_EncodeWithError(b *testing.B) {
	data := []byte("Hello, World!")
	encoder := &StdEncoder{Error: bytes.ErrTooLarge} // This will cause an error

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		encoder.Encode(data)
	}
}

// BenchmarkStdDecoder_DecodeWithError benchmarks the standard base58 decoder with invalid data
func BenchmarkStdDecoder_DecodeWithError(b *testing.B) {
	// Create invalid base58 data (contains characters not in the alphabet)
	invalidData := []byte("INVALID_BASE58_DATA!!!")
	decoder := NewStdDecoder()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		decoder.Decode(invalidData)
	}
}

// BenchmarkStdEncoder_EncodeBitcoinStyle benchmarks the standard base58 encoder with Bitcoin-style data
func BenchmarkStdEncoder_EncodeBitcoinStyle(b *testing.B) {
	// Create data that simulates Bitcoin addresses (32 bytes)
	data := make([]byte, 32)
	for i := range data {
		data[i] = byte((i*11 + 17) % 256)
	}
	encoder := NewStdEncoder()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		encoder.Encode(data)
	}
}

// BenchmarkStdDecoder_DecodeBitcoinStyle benchmarks the standard base58 decoder with Bitcoin-style data
func BenchmarkStdDecoder_DecodeBitcoinStyle(b *testing.B) {
	// Create encoded Bitcoin-style data for testing
	encoder := NewStdEncoder()
	data := make([]byte, 32)
	for i := range data {
		data[i] = byte((i*11 + 17) % 256)
	}
	encoded := encoder.Encode(data)
	decoder := NewStdDecoder()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		decoder.Decode(encoded)
	}
}

// BenchmarkStdEncoder_EncodeHashData benchmarks the standard base58 encoder with hash data
func BenchmarkStdEncoder_EncodeHashData(b *testing.B) {
	// Create data that simulates hash values (64 bytes)
	data := make([]byte, 64)
	for i := range data {
		data[i] = byte((i*13 + 19) % 256)
	}
	encoder := NewStdEncoder()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		encoder.Encode(data)
	}
}

// BenchmarkStdDecoder_DecodeHashData benchmarks the standard base58 decoder with hash data
func BenchmarkStdDecoder_DecodeHashData(b *testing.B) {
	// Create encoded hash data for testing
	encoder := NewStdEncoder()
	data := make([]byte, 64)
	for i := range data {
		data[i] = byte((i*13 + 19) % 256)
	}
	encoded := encoder.Encode(data)
	decoder := NewStdDecoder()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		decoder.Decode(encoded)
	}
}

// BenchmarkStreamingVsStandard benchmarks streaming vs standard encoding/decoding
func BenchmarkStreamingVsStandard(b *testing.B) {
	data := make([]byte, 10*1024) // 10KB
	for i := range data {
		data[i] = byte(i % 256)
	}

	b.Run("standard_encoder", func(b *testing.B) {
		b.ResetTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			NewStdEncoder().Encode(data)
		}
	})

	b.Run("streaming_encoder", func(b *testing.B) {
		b.ResetTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			var buf bytes.Buffer
			encoder := NewStreamEncoder(&buf)
			encoder.Write(data)
			encoder.Close()
		}
	})

	b.Run("standard_decoder", func(b *testing.B) {
		encoded := NewStdEncoder().Encode(data)
		b.ResetTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			NewStdDecoder().Decode(encoded)
		}
	})

	b.Run("streaming_decoder", func(b *testing.B) {
		encoded := NewStdEncoder().Encode(data)
		b.ResetTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			reader := mock.NewFile(encoded, "test.bin")
			decoder := NewStreamDecoder(reader)
			io.Copy(io.Discard, decoder)
		}
	})
}

// BenchmarkLargeFileStreaming benchmarks streaming performance for various data sizes
func BenchmarkLargeFileStreaming(b *testing.B) {
	sizes := []int{1024, 10 * 1024, 50 * 1024} // 1KB, 10KB, 50KB (reduced from 100KB and 1MB)

	for _, size := range sizes {
		data := make([]byte, size)
		for i := range data {
			data[i] = byte(i % 256)
		}

		b.Run(fmt.Sprintf("encode_%dKB", size/1024), func(b *testing.B) {
			b.ResetTimer()
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				var buf bytes.Buffer
				encoder := NewStreamEncoder(&buf)
				encoder.Write(data)
				encoder.Close()
			}
		})

		b.Run(fmt.Sprintf("decode_%dKB", size/1024), func(b *testing.B) {
			encoded := NewStdEncoder().Encode(data)
			b.ResetTimer()
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				reader := mock.NewFile(encoded, "test.bin")
				decoder := NewStreamDecoder(reader)
				io.Copy(io.Discard, decoder)
			}
		})
	}
}

// BenchmarkStreamingBufferSizes benchmarks streaming performance with different buffer sizes
func BenchmarkStreamingBufferSizes(b *testing.B) {
	data := make([]byte, 20*1024) // 20KB (reduced from 100KB)
	for i := range data {
		data[i] = byte(i % 256)
	}

	bufferSizes := []int{256, 512, 1024, 2048} // Reduced from 5 sizes to 4

	for _, bufSize := range bufferSizes {
		b.Run(fmt.Sprintf("buffer_%d", bufSize), func(b *testing.B) {
			b.ResetTimer()
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				var buf bytes.Buffer
				encoder := NewStreamEncoder(&buf)

				// Write in chunks of buffer size
				for j := 0; j < len(data); j += bufSize {
					end := j + bufSize
					if end > len(data) {
						end = len(data)
					}
					encoder.Write(data[j:end])
				}
				encoder.Close()
			}
		})
	}
}
