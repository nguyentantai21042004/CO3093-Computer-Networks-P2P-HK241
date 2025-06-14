# P2P File Sharing System

## Giới thiệu
Đây là dự án môn học Mạng Máy Tính (CO3093) tại Trường Đại học Bách Khoa - Đại học Quốc gia TP.HCM. Dự án này tập trung vào việc xây dựng một hệ thống chia sẻ file ngang hàng (P2P) sử dụng giao thức BitTorrent, giúp sinh viên hiểu rõ về cách hoạt động của mạng P2P và các giao thức mạng phổ biến.

## Tính năng chính
- Chia sẻ và tải file thông qua mạng P2P
- Hỗ trợ nhiều peer cùng lúc
- Tracker server để quản lý thông tin về các file và peer
- Giao thức BitTorrent để chia sẻ file hiệu quả
- Hỗ trợ tải file theo từng phần (chunk)

## Cấu trúc dự án
```
.
├── cmd/            # Chứa các file thực thi chính
├── pkg/            # Chứa các package chính của ứng dụng
│   ├── models/     # Định nghĩa các cấu trúc dữ liệu
│   ├── torrent/    # Xử lý torrent file và giao thức BitTorrent
│   └── utils/      # Các tiện ích hỗ trợ
├── tracker/        # Server tracker quản lý thông tin file và peer
└── config/         # Cấu hình ứng dụng
```

## Yêu cầu hệ thống
- Go 1.21.4 trở lên
- Các thư viện phụ thuộc được liệt kê trong file `go.mod`

## Cài đặt và chạy
1. Clone repository:
```bash
git clone [repository-url]
```

2. Cài đặt dependencies:
```bash
go mod download
```

3. Chạy tracker server:
```bash
go run cmd/main.go
```

## Cách hoạt động
1. **Tracker Server**: 
   - Quản lý thông tin về các file được chia sẻ
   - Theo dõi các peer đang online
   - Cung cấp thông tin về các peer có file cần tải

2. **Peer**:
   - Có thể vừa là người chia sẻ (seeder) vừa là người tải (leecher)
   - Tải file theo từng phần từ nhiều peer khác nhau
   - Tự động chia sẻ các phần đã tải được

## Học tập
Dự án này giúp sinh viên hiểu về:
- Kiến trúc mạng P2P
- Giao thức BitTorrent
- Xử lý đồng thời trong Go
- Quản lý kết nối mạng
- Chia sẻ tài nguyên trong mạng phân tán

## Đóng góp
Mọi đóng góp đều được hoan nghênh. Vui lòng tạo issue hoặc pull request để đóng góp vào dự án.

## Giấy phép
Dự án này được phát triển cho mục đích học tập và nghiên cứu. 