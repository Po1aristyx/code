#include <iostream>
#include <string>
#include <vector>
#include <fstream>

using namespace std;

// 用户类
class User {
public:
    // 用户注册
    static void registerUser(string name, string phone, string password) {
        // 将用户信息保存到文件中
        ofstream file("users.txt", ios::app);
        file << name << "," << phone << "," << password << endl;
        file.close();
        cout << "注册成功!" << endl;
    }

    // 用户登录
    static bool login(string phone, string password) {
        // 从文件中读取用户信息并验证
        ifstream file("users.txt");
        string line;
        while (getline(file, line)) {
            size_t pos = line.find(",");
            string name = line.substr(0, pos);
            line = line.substr(pos + 1);
            pos = line.find(",");
            string userPhone = line.substr(0, pos);
            string userPassword = line.substr(pos + 1);
            if (userPhone == phone && userPassword == password) {
                file.close();
                return true;
            }
        }
        file.close();
        return false;
    }
};

// 预约类
class Reservation {
public:
    // 预约餐位
    static void makeReservation(string phone, string time, string tableType) {
        // 检查餐位是否可用
        if (isTableAvailable(tableType)) {
            // 生成预约单
            string reservationId = generateReservationId();
            cout << "预约成功,您的预约单号为: " << reservationId << endl;
            cout << "预约时间: " << time << ", 餐桌类型: " << tableType << endl;
            // 将预约信息保存到文件中
            ofstream file("reservations.txt", ios::app);
            file << reservationId << "," << phone << "," << time << "," << tableType << endl;
            file.close();
        }
        else {
            cout << "很抱歉,该时段的" << tableType << "餐位已经被预订完了。" << endl;
        }
    }

private:
    // 检查餐位是否可用
    static bool isTableAvailable(string tableType) {
        // 这里可以实现具体的餐位检查逻辑
        return true;
    }

    // 生成预约单号
    static string generateReservationId() {
        // 这里可以实现具体的预约单号生成逻辑
        return "R001";
    }
};

int main() {
    // 用户登录
    string phone, password;
    cout << "请输入手机号: ";
    cin >> phone;
    cout << "请输入密码: ";
    cin >> password;
    if (User::login(phone, password)) {
        cout << "登录成功!" << endl;

        // 预约餐位
        string time, tableType;
        cout << "请输入预约时间(格式:HH:MM): ";
        cin >> time;
        cout << "请输入餐桌类型(如:包厢、四人桌、二人桌): ";
        cin >> tableType;
        Reservation::makeReservation(phone, time, tableType);
    }
    else {
        cout << "登录失败,请检查手机号和密码是否正确。" << endl;
    }

    return 0;
}