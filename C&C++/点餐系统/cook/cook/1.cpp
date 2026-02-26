#include <iostream>
#include <vector>
#include <string>
#include <fstream>
#include <ctime>
#include <iomanip>
#include <sstream>

using namespace std;


class User {
protected:
    string username;
    string phoneNumber;
    string password;
public:
    
    User(string uname, string phone, string pwd)
        : username(uname), phoneNumber(phone), password(pwd) {}

    
    virtual void registerUser() = 0;
    virtual bool login(string uname, string pwd) = 0;
    string getUsername() {
        return username;
    }

    string getPhoneNumber() {
        return phoneNumber;
    }

    void setPhoneNumber(string phone) {
        phoneNumber = phone;
    }

    void setPassword(string pwd) {
        password = pwd;
    }
};

class Customer : public User {
private:
    struct Reservation {
        int reservationID;           // 预定ID
        string reservationTime; // 预定时间
        string seatType;        // 座位类型
        int seatNumber;         // 座位号
        vector<string> foodItems; // 预定的餐食列表
        double totalPrice;       // 总价
        double discount;         // 折扣
        double deposit;          // 定金
    };
    vector<Reservation> reservations;
public:
    Customer(string uname, string phone, string pwd)
        : User(uname, phone, pwd) {}

    void registerUser() override {
        ofstream file("customers.txt", ios::app);
        if (file.is_open()) {
            if (phoneNumber.length() != 11) {
                cout << "错误：手机号必须为11位。请重新输入。" << endl;
                return;
            }
            file << username << " " << phoneNumber << " " << password << endl;
            file.close();
            cout << "注册成功" << endl;
        }
        else {
            cout << "无法打开文件进行注册." << endl;
        }
    }

    bool login(string uname, string pwd) override {
        ifstream file("customers.txt");
        string u, p, phone;
        if (file.is_open()) {
            while (file >> u >> phone >> p) { 
                if (u == uname && p == pwd) {
                    username = uname;
                    phoneNumber = phone;
                    password = pwd;
                    cout << "登录成功" << endl;
                    return true;
                }
            }
            file.close();
        }
        cout << "登录失败" << endl;
        return false;
    }
    void displayFoodMenu() {
        cout << "菜单:" << endl;
        cout << "1.汉堡" << endl;
        cout << "2.比萨" << endl;
        cout << "3.沙拉" << endl;
        cout << "4.苏打水" << endl;
        cout << "5.咖啡" << endl;
        cout << endl;
    }

    void bookSeat() {
        int reservationID = time(0);
        string reservationTime;
        string seatType;
        int seatNumber;
        vector<string> foodItems;
        double totalPrice = 0.0;
        double discount = 0.9;
        double deposit;

        cout << "请输入预定时间 (YYYY-MM-DD HH:MM): ";
        cin.ignore();
        getline(cin, reservationTime);
        cout << "输入座位种类(包厢,四人桌,二人桌): ";
        cin >> seatType;

        cout << "可用座位号：";
        if (seatType == "包厢") {
            cout << "001、002、003、...、050" << endl;
        }
        else if (seatType == "四人桌") {
            cout << "001、002、003、...、025" << endl;
        }
        else if (seatType == "二人桌") {
            cout << "001、002、003、...、050" << endl;
        }
        else {
            cout << "无效的座位类型" << endl;
            return;
        }

        cout << "请输入座位号：";
        cin >> seatNumber;

        bool seatAvailable = true;
        for (const auto& res : reservations) {
            if (res.seatType == seatType && res.seatNumber == seatNumber) {
                seatAvailable = false;
                break;
            }
        }

        if (!seatAvailable) {
            cout << seatType << " 中的座位 " << seatNumber << " 已经被预订。" << endl;
            return;
        }


        displayFoodMenu();
        cout << "请输入食物数量: ";
        int foodCount;
        cin >> foodCount;
        cin.ignore();
        for (int i = 0; i < foodCount; ++i) {
            int foodChoice;
            string foodItem;
            double price;
            cout << "请输入食物数量" << i + 1 << "请输入编号: ";
            cin >> foodChoice;
            cin.ignore();
            switch (foodChoice) {
            case 1:
                foodItem = "汉堡";
                price = 8;
                break;
            case 2:
                foodItem = "比萨";
                price = 25;
                break;
            case 3:
                foodItem = "沙拉";
                price = 5;
                break;
            case 4:
                foodItem = "苏打水";
                price = 7;
                break;
            case 5:
                foodItem = "咖啡";
                price = 10;
                break;
            default:
                cout << "无效选择" << endl;
                i--;
                continue;
            }
            foodItems.push_back(foodItem);
            totalPrice += price;
        }

        totalPrice *= discount;
        deposit = totalPrice * 0.3;

        reservations.push_back({ reservationID, reservationTime, seatType,seatNumber, foodItems, totalPrice, discount, deposit });
        cout << fixed << setprecision(2);
        cout << " 预定座位和餐食成功。预定ID: " << reservationID << endl;
        cout << " 折扣后总价: " << totalPrice << endl;
        cout << " 支付定金: " << deposit << endl;
    }

    void cancelBooking() {
        int reservationID;
        cout << "请输入要取消的预订ID: ";
        cin >> reservationID;

        time_t now = time(0);
        tm ltm = {};
        localtime_s(&ltm, &now);

        for (auto it = reservations.begin(); it != reservations.end(); ++it) {
            if (it->reservationID == reservationID) { 
                tm reservationTime = {};
                istringstream ss(it->reservationTime);
                ss >> get_time(&reservationTime, "%Y-%m-%d %H:%M");

                double hoursDiff = difftime(mktime(&reservationTime), mktime(&ltm)) / 3600.0;
                if (hoursDiff < 3) {
                    cout << "预订时间已不到三小时，无法取消" << endl;
                    return;
                }

                reservations.erase(it);
                cout << "您的预订已成功取消" << endl;
                return;
            }
        }
        cout << "您的预订ID未找到." << endl;
    }

    void modifyBooking() {
        int reservationID;
        cout << "请输入要修改的预订ID: ";
        cin >> reservationID;

        time_t now = time(0);
        tm ltm = {};
        localtime_s(&ltm, &now);

        for (auto& res : reservations) {
            if (res.reservationID == reservationID) {
                tm reservationTime = {};
                istringstream ss(res.reservationTime);
                ss >> get_time(&reservationTime, "%Y-%m-%d %H:%M");

                double hoursDiff = difftime(mktime(&reservationTime), mktime(&ltm)) / 3600.0;
                if (hoursDiff < 3) {
                    cout << "预订时间已不到三小时，无法修改" << endl;
                    return;
                }

                string newSeatType;
                cout << "请输入新的座位种类: ";
                cin >> newSeatType;
                res.seatType = newSeatType;
                cout << "您的预订已成功修改" << endl;
                return;
            }
        }
        cout << "您的预订ID未找到" << endl;
    }

    void saveReservations() {
        ofstream file("reservations.txt");
        if (file.is_open()) {
            for (const auto& res : reservations) {
                file << res.reservationID << " " << res.reservationTime << " " << res.seatType << " ";
                for (const auto& item : res.foodItems) {
                    file << item << ",";
                }
                file << " " << res.totalPrice << " " << res.discount << " " << res.deposit << endl;
            }
            file.close();
        }
    }

    
    void loadReservations() {
        ifstream file("reservations.txt");
        if (file.is_open()) {
            string line;
            while (getline(file, line)) {
                istringstream iss(line);
                Reservation res;
                string foodItems;
                if (iss >> res.reservationID >> res.reservationTime >> res.seatType >> foodItems >> res.totalPrice >> res.discount >> res.deposit) {
                    stringstream ss(foodItems);
                    string item;
                    while (getline(ss, item, ',')) {
                        if (!item.empty()) {
                            res.foodItems.push_back(item);
                        }
                    }
                    reservations.push_back(res);
                }
            }
            file.close();
        }
    }
};


class Admin : public User {
public:
    
    Admin(string uname, string phone, string pwd)
        : User(uname, phone, pwd) {}

    void registerUser() override {
        ofstream file("admins.txt", ios::app);
        if (file.is_open()) {
            if (phoneNumber.length() != 11) {
                cout << "错误：管理员手机号必须为11位。请重新输入。" << endl;
                return;
            }
            file << username << " " << phoneNumber << " " << password << endl;
            file.close();
            cout << "注册成功" << endl;
        }
        else {
            cout << "无法打开文件进行注册." << endl;
        }
    }

    bool login(string uname, string pwd) override {
        ifstream file("admins.txt"); 
        string u, p, phone;
        if (file.is_open()) {
            while (file >> u >> phone >> p) {
                if (u == uname && p == pwd) {
                    username = uname;
                    phoneNumber = phone;
                    password = pwd;
                    cout << "登录成功" << endl;
                    return true;
                }
            }
            file.close();
        }
        cout << "登录失败" << endl;
        return false;
    }

    void deleteUser(string uname) {
        ifstream inputFile("customers.txt");
        ofstream tempFile("temp.txt");

        if (inputFile.is_open() && tempFile.is_open()) {
            string u, p, phone;
            while (inputFile >> u >> phone >> p) {
                if (u != uname) {
                    tempFile << u << " " << phone << " " << p << endl;
                }
            }
            inputFile.close();
            tempFile.close();
            remove("customers.txt");
            rename("temp.txt", "customers.txt");
            cout << "用户删除成功" << endl;
        }
    }

    void deleteReservation(int reservationID) {
        ifstream inputFile("reservations.txt");
        ofstream tempFile("temp.txt");

        if (inputFile.is_open() && tempFile.is_open()) {
            string line;
            while (getline(inputFile, line)) {
                istringstream iss(line);
                int resID;
                iss >> resID;
                if (resID != reservationID) {
                    tempFile << line << endl;
                }
            }
            inputFile.close();
            tempFile.close();
            remove("reservations.txt");
            rename("temp.txt", "reservations.txt");
            cout << "预定删除成功" << endl;
        }
    }

    void viewUsers() {
        ifstream file("customers.txt");
        if (file.is_open()) {
            string u, phone, p;
            while (file >> u >> phone >> p) {
                cout << " 用户名: " << u << ", 电话: " << phone << endl;
            }
            file.close();
        }
    }

    void viewReservations() {
        ifstream file("reservations.txt");
        if (file.is_open()) {
            string line;
            while (getline(file, line)) {
                cout << line << endl;
            }
            file.close();
        }
    }
};

int main() {
    Customer customer("default", "00000000000", "password");
    Admin admin("admin", "11111111111", "admin");

    customer.loadReservations();

    int choice;
    do {
        cout << endl;
        cout << "****************************************************" << endl;
        cout << "1.以顾客身份注册" << endl;
        cout << "2.以管理员身份注册" << endl;
        cout << "3. 顾客登录" << endl;
        cout << "4.管理员登录" << endl;
        cout << "0.退出" << endl;
        cout << "****************************************************" << endl;
        cout << endl;
        cout << "输入选择: ";
        cin >> choice;

        switch (choice) {
        case 1: {
            string uname, phone, pwd;
            cout << "输入用户名: ";
            cin >> uname;
            cout << "输入电话号码: ";
            cin >> phone;
            cout << "输入密码: ";
            cin >> pwd;
            cout << endl;
            Customer newCustomer(uname, phone, pwd);
            newCustomer.registerUser();
            break;
        }
        case 2: {
            string uname, phone, pwd;
            cout << "输入管理员用户名: ";
            cin >> uname;
            cout << "输入管理员电话号码: ";
            cin >> phone;
            cout << "输入管理员密码: ";
            cin >> pwd;
            cout << endl;
            Admin newAdmin(uname, phone, pwd);
            newAdmin.registerUser();
            break;
        }
        case 3: {
            string uname, pwd;
            cout << "输入用户名: ";
            cin >> uname;
            cout << "输入密码: ";
            cin >> pwd;
            if (customer.login(uname, pwd)) {
                int custChoice;
                do {
                    cout << "1. 预定座位" << endl;
                    cout << "2. 取消预定" << endl;
                    cout << "3. 修改预定" << endl;
                    cout << "0. 退出" << endl;
                    cout << "输入选择: ";
                    cin >> custChoice;


                    switch (custChoice) {
                    case 1:
                        customer.bookSeat();
                        customer.saveReservations();
                        break;
                    case 2:
                        customer.cancelBooking();
                        customer.saveReservations();
                        break;
                    case 3:
                        customer.modifyBooking();
                        customer.saveReservations();
                        break;
                    case 0:
                        break;
                    default:
                        cout << "无效选择" << endl;
                    }
                } while (custChoice != 0);
            }
            break;
        }
        case 4: {
            string uname, pwd;
            cout << "输入管理员用户名: ";
            cin >> uname;
            cout << "输入管理员密码: ";
            cin >> pwd;
            if (admin.login(uname, pwd)) {
                int adminChoice;
                do {
                    cout << "1.查看用户" << endl;
                    cout << "2.查看预定" << endl;
                    cout << "3.删除用户" << endl;
                    cout << "4.删除预定" << endl;
                    cout << "0.退出" << endl;
                    cout << endl;
                    cout << "输入选择: ";
                    cin >> adminChoice;

                    switch (adminChoice) {
                    case 1:
                        admin.viewUsers();
                        break;
                    case 2:
                        admin.viewReservations();
                        break;
                    case 3: {
                        string uname;
                        cout << "输入要删除的用户名: ";
                        cin >> uname;
                        admin.deleteUser(uname);
                        break;
                    }
                    case 4: {
                        int reservationID;
                        cout << "输入要删除的预订ID: ";
                        cin >> reservationID;
                        admin.deleteReservation(reservationID);
                        break;
                    }
                    case 5: {
                        int reservationID;
                        cout << "输入要修改的预订ID: ";
                        cin >> reservationID;
                        admin.deleteReservation(reservationID);
                        break;
                    }
                    case 0:
                        break;
                    default:
                        cout << "无效选择" << endl;
                    }
                } while (adminChoice != 0);
            }
            break;
        }
        case 0:
            break;
        default:
            cout << "无效选择" << endl;
        }
    } while (choice != 0);

    return 0;
}