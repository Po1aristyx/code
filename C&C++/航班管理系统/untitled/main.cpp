#include <iostream>
#include <string>
#include <vector>
#include <map>
#include <ctime>
#include <algorithm>
#include <fstream>
#include <sstream>
using namespace std;
// 客户类
struct Customer {
    string name;
    int ticket_count;
    int cabin_class; // 1, 2, or 3

    // 序列化客户信息
    string serialize() const {
        return name + "," + to_string(ticket_count) + "," + to_string(cabin_class);
    }

    // 反序列化客户信息
    static Customer deserialize(const string& data) {
        Customer customer;
        stringstream ss(data);
        getline(ss, customer.name, ',');
        string temp;
        getline(ss, temp, ',');
        customer.ticket_count = stoi(temp);
        getline(ss, temp, ',');
        customer.cabin_class = stoi(temp);
        return customer;
    }
};

// 航班类
struct Flight {
    string flight_number;
    string airplane_number;
    string day_of_week; // e.g., "Monday"
    int crew_quota;
    int total_tickets;
    vector<Customer> booked_customers;

    // 计算余票
    int remaining_tickets() const {
        int booked = 0;
        for (const auto& customer : booked_customers) {
            booked += customer.ticket_count;
        }
        return total_tickets - booked;
    }

    // 序列化航班信息
    string serialize() const {
        string data = flight_number + "," + airplane_number + "," + day_of_week + "," +
                      to_string(crew_quota) + "," + to_string(total_tickets) + "\n";
        data += to_string(booked_customers.size()) + "\n";
        for (const auto& customer : booked_customers) {
            data += customer.serialize() + "\n";
        }
        return data;
    }

    // 反序列化航班信息
    static Flight deserialize(ifstream& infile) {
        Flight flight;
        string line;

        // 读取航班基本信息
        getline(infile, line);
        stringstream ss(line);
        getline(ss, flight.flight_number, ',');
        getline(ss, flight.airplane_number, ',');
        getline(ss, flight.day_of_week, ',');
        getline(ss, line, ',');
        flight.crew_quota = stoi(line);
        getline(ss, line, ',');
        flight.total_tickets = stoi(line);

        // 读取已订票的客户数量
        int customer_count = 0;
        getline(infile, line);
        customer_count = stoi(line);

        // 读取每个客户的信息
        for (int i = 0; i < customer_count; ++i) {
            getline(infile, line);
            flight.booked_customers.push_back(Customer::deserialize(line));
        }

        return flight;
    }
};

// 航班系统类
class FlightSystem {
private:
    map<string, Flight> flights; // Key: flight_number

    // 将星期几转换为整数（0=Monday, 6=Sunday）
    int day_to_int(const string& day) const {
        vector<string> days = { "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday" };
        auto it = find(days.begin(), days.end(), day);
        if (it != days.end()) {
            return distance(days.begin(), it);
        }
        return -1; // Invalid day
    }

    // 获取下一个航班日期
    string get_next_flight_date(const string& day_of_week) const {
        time_t t = time(nullptr);
        tm* now = localtime(&t);
        int current_weekday = now->tm_wday; // Sunday=0, Monday=1, ..., Saturday=6
        // 调整为0=Monday, ..., 6=Sunday
        current_weekday = (current_weekday == 0) ? 6 : current_weekday - 1;

        int flight_weekday = day_to_int(day_of_week);
        if (flight_weekday == -1) return "Invalid day";

        int days_ahead = (flight_weekday - current_weekday + 7) % 7;
        if (days_ahead == 0) days_ahead = 7; // Next week

        tm next_flight = *now;
        next_flight.tm_mday += days_ahead;
        mktime(&next_flight); // Normalize the time structure

        char buffer[11];
        strftime(buffer, sizeof(buffer), "%Y-%m-%d", &next_flight);
        return string(buffer);
    }

public:
    // 加载航班数据
    void load_data(const string& filename) {
        ifstream infile(filename);
        if (!infile.is_open()) {
            cout << "无法打开数据文件，开始使用空数据。" << endl;
            return;
        }

        string line;
        while (getline(infile, line)) {
            if (line.empty()) continue;
            stringstream ss(line);
            // 回到文件流开始一行读取
            infile.seekg(-static_cast<int>(line.length()) - 1, ios_base::cur);
            Flight flight = Flight::deserialize(infile);
            flights[flight.flight_number] = flight;
        }

        infile.close();
    }

    // 保存航班数据
    void save_data(const string& filename) const {
        ofstream outfile(filename);
        if (!outfile.is_open()) {
            cout << "无法打开数据文件进行保存。" << endl;
            return;
        }

        for (const auto& pair : flights) {
            outfile << pair.second.serialize();
        }

        outfile.close();
    }

    // 添加航班
    void add_flight(const Flight& flight) {
        if (flights.find(flight.flight_number) != flights.end()) {
            cout << "航班号已存在，无法添加。" << endl;
            return;
        }
        flights[flight.flight_number] = flight;
        cout << "航班添加成功。" << endl;
    }

    // 查询航线
    void query_routes(const string& destination) const {
        vector<Flight> results;
        for (const auto& pair : flights) {
            const Flight& flight = pair.second;
            // 简化：假设目的地与航班号后缀关联
            if (flight.flight_number.size() >= 2 &&
                flight.flight_number.substr(flight.flight_number.size() - 2) == destination) {
                results.push_back(flight);
            }
        }

        if (!results.empty()) {
            cout << "找到以下直达航班：" << endl;
            for (const auto& flight : results) {
                string next_flight_date = get_next_flight_date(flight.day_of_week);
                cout << "航班号: " << flight.flight_number
                     << ", 飞机号: " << flight.airplane_number
                     << ", 星期: " << flight.day_of_week
                     << ", 最近航班日期: " << next_flight_date
                     << ", 余票: " << flight.remaining_tickets() << endl;
            }
        }
        else {
            cout << "没有直达航班，建议换乘方案：" << endl;
            // 简化示例，仅输出一个假设的换乘方案
            cout << "从A到B需在C中转，总时间最短。" << endl;
        }
    }

    // 订票
    void book_ticket(const string& flight_number, const Customer& customer) {
        auto it = flights.find(flight_number);
        if (it != flights.end()) {
            Flight& flight = it->second;
            if (flight.remaining_tickets() >= customer.ticket_count) {
                flight.booked_customers.push_back(customer);
                cout << "订票成功，座位号: " << flight.booked_customers.size() << endl;
            }
            else {
                cout << "余票不足，无法完成订票。" << endl;
            }
        }
        else {
            cout << "航班号不存在。" << endl;
        }
    }

    // 退票
    void refund_ticket(const string& flight_number, const string& customer_name) {
        auto it = flights.find(flight_number);
        if (it != flights.end()) {
            Flight& flight = it->second;
            auto& customers = flight.booked_customers;
            auto cust_it = find_if(customers.begin(), customers.end(),
                                   [&](const Customer& c) { return c.name == customer_name; });
            if (cust_it != customers.end()) {
                customers.erase(cust_it);
                cout << "退票成功。" << endl;
            }
            else {
                cout << "未找到该客户的订票记录。" << endl;
            }
        }
        else {
            cout << "航班号不存在。" << endl;
        }
    }

    // 显示所有航班（用于调试）
    void display_all_flights() const {
        for (const auto& pair : flights) {
            const Flight& flight = pair.second;
            cout << "航班号: " << flight.flight_number
                 << ", 飞机号: " << flight.airplane_number
                 << ", 星期: " << flight.day_of_week
                 << ", 乘员定额: " << flight.crew_quota
                 << ", 总票量: " << flight.total_tickets
                 << ", 已订票客户数: " << flight.booked_customers.size() << endl;
        }
    }
};

// 显示主菜单
void display_menu() {
    cout << "\n===== 航班订票系统 =====" << endl;
    cout << "1. 查询航线" << endl;
    cout << "2. 订票" << endl;
    cout << "3. 退票" << endl;
    cout << "4. 添加航班" << endl;
    cout << "5. 显示所有航班" << endl;
    cout << "退出请输入Q" << endl;
    cout << "请选择操作 (1-6): ";
}

int main() {
    FlightSystem system;
    const string data_file = "flights_data.txt";

    // 加载航班数据
    system.load_data(data_file);

    char choice;
    while (true) {
        display_menu();
        cin >> choice;
        cin.ignore(); // 清除输入缓冲区

        switch (choice) {
            case '1': { // 查询航线
                string destination;
                cout << "请输入目的地代码（如 BJ 表示北京）: ";
                cin >> destination;
                system.query_routes(destination);
                break;
            }
            case '2': { // 订票
                string flight_number, customer_name;
                int ticket_count, cabin_class;
                cout << "请输入航班号: ";
                cin >> flight_number;
                cout << "请输入客户姓名: ";
                cin >> customer_name;
                cout << "请输入订票数量: ";
                cin >> ticket_count;
                cout << "请输入舱位等级（1, 2 或 3）: ";
                cin >> cabin_class;

                Customer customer = { customer_name, ticket_count, cabin_class };
                system.book_ticket(flight_number, customer);
                break;
            }
            case '3': { // 退票
                string flight_number, customer_name;
                cout << "请输入航班号: ";
                cin >> flight_number;
                cout << "请输入客户姓名: ";
                cin >> customer_name;
                system.refund_ticket(flight_number, customer_name);
                break;
            }
            case '4': { // 添加航班
                Flight flight;
                cout << "请输入航班号: ";
                cin >> flight.flight_number;
                cout << "请输入飞机号: ";
                cin >> flight.airplane_number;
                cout << "请输入星期几（如 Monday）: ";
                cin >> flight.day_of_week;
                cout << "请输入乘员定额: ";
                cin >> flight.crew_quota;
                cout << "请输入总票量: ";
                cin >> flight.total_tickets;
                flight.booked_customers.clear(); // 新航班没有已订票客户

                system.add_flight(flight);
                break;
            }
            case '5': { // 显示所有航班
                system.display_all_flights();
                break;
            }
            case 'Q': { // 退出
                system.save_data(data_file);
                cout << "数据已保存，退出系统。" << endl;
                return 0;
            }
            default:
                cout << "无效的选择，请重新输入。" << endl;
                break;
        }
    }

    return 0;
}