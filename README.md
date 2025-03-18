# Clean Code
## SOLID
SOLID là một tập hợp các nguyên tắc thiết kế phần mềm trong lập trình hướng đối tượng được đưa ra bởi Robert C. Martin, giúp các nhà phát triển tạo ra phần mềm dễ bảo trì và mở rộng. Mặc dù Go không phải là ngôn ngữ hoàn toàn hướng đối tượng, nhưng nguyên tắc này vẫn có thể áp dụng hiệu quả trong việc thiết kế code với Go.
### Single Responsibility Principle
> A module should have only one reason to change.

- Nội dung của nguyên lý này chỉ ra rằng mỗi package, struct và function đều nên chỉ giữ một trách nhiệm duy nhất. Nói cách khác, package, struct, function chỉ nên có một lý do duy nhất để thay đổi.

#### Struct
- Một struct có quá nhiều chức năng cũng sẽ trở nên cồng kềnh và phức tạp, việc này dẫn đến việc thay đổi code sẽ rất khó khăn, mất nhiều thời gian, còn dễ gây ảnh hưởng tới các module đang hoạt động khác.
- Ví dụ về struct vi phạm nguyên tắc SRP.
    ```go
    // Define user model
    type User struct {
	    ID        uint
	    FirstName string
	    LastName  string
	    Email     string
	    Password  string
    }

    func (u *User) SaveToDatabase() {
	    // Save user info in database
        // ...
    }

    func (*User) Authenticate(email, password string) {
	    // Authenticate user logic
        // ...
    }
    ```
    - Ở ví dụ trên, struct `User` vừa chịu trách nhiệm định nghĩa thực thể người dùng, vừa chịu trách nhiệm tương tác với database để lưu thông tin người dùng, vừa thực hiện các logic về việc xác thực người dùng. Điều này đã vi phạm nghiêm trọng nguyên tắc SRP khi struct `User` chịu một lúc 3 trách nhiệm.
    - Để đảm bảo SRP, ta có thể tách struct trên thành 3 phần riêng biệt.
    - Struct `User` trong `package model` để định nghĩa cấu trúc của thực thể người dùng
    ```go
    package model

    type User struct {
	    ID        uint
	    FirstName string
	    LastName  string
	    Email     string
	    Password  string
    }
    ```
    - Struct `UserRepository` trong `package repository` để thực hiện tương tác với database
    ```go
    package repository

    type UserRepository struct {
	    db     *gorm.DB
	    logger *zap.Logger
    }

    func (u *userRepository) CreateUser(ctx context.Context, user model.User) error {
        // Create new user to database
    }

    func (u *userRepository) UpdateUser(ctx context.Context, user model.User) error {
        // Update user in database
    }
    ```
    - Struct `UserService` trong `package service` để thực hiện các business logic
    ```go
    package service

    type userService struct {
	    userRepo repository.UserRepository
	    jwt      util.JwtUtil
	    logger   *zap.Logger
    }

    func (u *userService) Authenticate(ctx context.Context, email, password string) (accessToken string,err error) {
        // Authentication logic
    }
    ```

#### Function
- Các function trong Go cũng nên tuân thủ nguyên tắc SRP bằng cách chỉ nên tập trung xử lý một tác vụ duy nhất. Một function xử lí quá nhiều vấn đề có thể trở nên khó hiểu, khó bảo trì và kiểm thử. Bằng cách chia các hàm phức tạp thành các hàm nhỏ hơn và tập trung xử lý một vấn đề duy nhất, ta có thể làm cho code trở nên dễ hiểu, dễ bảo trì và kiểm thử.

### Open/Closed Principle
> Software entities (classes, modules, functions, etc...) should be open for extension, but closed for modification.

- 

### Liskov Substitution Principle
>Objects in a program should be replaceable with instances of their subtypes without altering the correctness of that program.

- Nội dung của nguyên lý này chỉ ra rằng đối tượng của lớp cha có thể thay thế bằng các lớp con của nó mà không làm thay đổi tính đúng đắn của chương trình.
- Trong Go, nguyên lý này được thể hiện thông qua `interface`.
- Nếu một kiểu `T` implement interface `I` thì các đối tượng của kiểu `T` có thể thay thế cho kiểu `I` mà không làm thay đổi tính chất, hành vi mong đợi của chương trình (tính đúng đắn, nhiệm vụ cần thực hiện,...). Nghĩa là bất kỳ một kiểu (type) nào implement interface `I` thì hành vi của các phương thức của nó phải nhất quán với kỳ vọng được đặt ra bởi interface đó.

### Interface Segregation








