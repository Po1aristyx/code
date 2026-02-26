#include <iostream>
#include <vector>
#include <string>
#include <fstream>
#include <ctime>

// 用户基类，包含用户名、手机号和密码
class User {
protected:
    std::string username;
    std::string phoneNumber;
    std::string password;
public:
    // 构造函数初始化用户信息
    User(std::string uname, std::string phone, std::string pwd)
        : username(uname), phoneNumber(phone), password(pwd) {}

    // 纯虚函数，供子类实现注册功能
    virtual void registerUser() = 0;
    // 纯虚函数，供子类实现登录功能
    virtual bool login(std::string uname, std::string pwd) = 0;
};

// 顾客类，继承自User，包含预定信息
class Customer : public User {
private:
    struct Reservation {
        int reservationID;           // 预定ID
        std::string reservationTime; // 预定时间
        std::string seatType;        // 座位类型
    };
    std::vector<Reservation> reservations; // 顾客的预定列表
public:
    // 构造函数，初始化顾客信息
    Customer(std::string uname, std::string phone, std::string pwd)
        : User(uname, phone, pwd) {}

    // 实现注册功能，将用户信息保存到文件
    void registerUser() override {
        std::ofstream file("customers.txt", std::ios::app); // 以追加方式打开文件
        if (file.is_open()) {
            file << username << " " << phoneNumber << " " << password << std::endl;
            file.close();
            std::cout << "Registration successful.\n";
        }
        else {
            std::cout << "Unable to open file for registration.\n";
        }
    }

    // 实现登录功能，从文件读取用户信息并验证
    bool login(std::string uname, std::string pwd) override {
        std::ifstream file("customers.txt"); // 打开文件读取用户信息
        std::string u, p, phone;
        if (file.is_open()) {
            while (file >> u >> phone >> p) { // 逐行读取文件内容
                if (u == uname && p == pwd) { // 验证用户名和密码
                    username = uname;
                    phoneNumber = phone;
                    password = pwd;
                    std::cout << "Login successful.\n";
                    return true;
                }
            }
            file.close();
        }
        std::cout << "Login failed.\n";
        return false;
    }

    // 预定座位，记录预定信息
    void bookSeat() {
        int reservationID = std::time(0); // 使用当前时间生成预定ID
        std::string reservationTime;
        std::string seatType;

        // 获取用户输入的预定时间和座位类型
        std::cout << "Enter reservation time (YYYY-MM-DD HH:MM): ";
        std::cin.ignore();
        std::getline(std::cin, reservationTime);
        std::cout << "Enter seat type (e.g., VIP, Regular): ";
        std::cin >> seatType;

        reservations.push_back({ reservationID, reservationTime, seatType }); // 记录预定信息
        std::cout << "Seat booked successfully. Reservation ID: " << reservationID << "\n";
    }

    // 取消预定，根据预定ID删除预定信息
    void cancelBooking() {
        int reservationID;
        std::cout << "Enter reservation ID to cancel: ";
        std::cin >> reservationID;

        for (auto it = reservations.begin(); it != reservations.end(); ++it) {
            if (it->reservationID == reservationID) { // 查找对应预定ID
                reservations.erase(it); // 删除预定信息
                std::cout << "Booking cancelled successfully.\n";
                return;
            }
        }
        std::cout << "Reservation ID not found.\n";
    }

    // 修改预定，根据预定ID修改预定信息
    void modifyBooking() {
        int reservationID;
        std::cout << "Enter reservation ID to modify: ";
        std::cin >> reservationID;

        for (auto& res : reservations) {
            if (res.reservationID == reservationID) { // 查找对应预定ID
                std::string newSeatType;
                std::cout << "Enter new seat type: ";
                std::cin >> newSeatType;
                res.seatType = newSeatType; // 修改座位类型
                std::cout << "Booking modified successfully.\n";
                return;
            }
        }
        std::cout << "Reservation ID not found.\n";
    }
};

int main() {
    Customer cust("default", "0000000000", "default"); // 创建顾客对象，默认信息将被替换

    int choice;
    do {
        // 显示菜单选项
        std::cout << "1. Register\n2. Login\n3. Book Seat\n4. Cancel Booking\n5. Modify Booking\n6. Exit\n";
        std::cout << "Enter choice: ";
        std::cin >> choice;

        switch (choice) {
        case 1:
            // 注册新用户
        {
            std::string uname, phone, pwd;
            std::cout << "Enter username: ";
            std::cin >> uname;
            std::cout << "Enter phone number: ";
            std::cin >> phone;
            std::cout << "Enter password: ";
            std::cin >> pwd;
            Customer newCustomer(uname, phone, pwd);
            newCustomer.registerUser();
        }
        break;
        case 2:
        {
            // 用户登录
            std::string uname, pwd;
            std::cout << "Enter username: ";
            std::cin >> uname;
            std::cout << "Enter password: ";
            std::cin >> pwd;
            if (cust.login(uname, pwd)) {
                cust = Customer(uname, cust.phoneNumber, pwd); // 更新当前顾客对象
            }
        }
        break;
        case 3:
            // 预定座位
            if (!cust.login("dummy", "dummy")) {
                std::cout << "Please login first.\n";
            }
            else {
                cust.bookSeat();
            }
            break;
        case 4:
            // 取消预定
            if (!cust.login("dummy", "dummy")) {
                std::cout << "Please login first.\n";
            }
            else {
                cust.cancelBooking();
            }
            break;
        case 5:
            // 修改预定
            if (!cust.login("dummy", "dummy")) {
                std::cout << "Please login first.\n";
            }
            else {
                cust.modifyBooking();
            }
            break;
        case 6:
            std::cout << "Exiting...\n";
            break;
        default:
            std::cout << "Invalid choice. Try again.\n";
        }
    } while (choice != 6);

    return 0;
}
