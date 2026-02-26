#include <bits/stdc++.h>
#include <unordered_map>
#include <vector>
#include <algorithm>
#include <string>

using namespace std;

struct Customer {
    string name;
    int ticket_count;
    int cabin_class;
};

struct Flight {
    string flight_number;
    string airplane_number;
    string start;
    string end;
    string day_of_week;
    string date;
    string departure_time;
    string arrival_time;
    int crew_quota;
    int total_tickets;
    vector<Customer> booked_customers;
    vector<Customer> waitlist;

    Flight() : waitlist() {}

    int remaining_tickets() const {
        int available_tickets = total_tickets - crew_quota; // 先减去乘员定额
        int booked_tickets = 0;
        for (const auto& customer : booked_customers) {
            booked_tickets += customer.ticket_count;
        }
        return available_tickets - booked_tickets;
    }
    int flight_duration() const {
        auto to_minutes = [](const string& time) -> int {
            if (time.length() < 5) return 0;
            int hours = stoi(time.substr(0, 2));
            int minutes = stoi(time.substr(3, 2));
            return hours * 60 + minutes;
        };
        int departure = to_minutes(departure_time);
        int arrival = to_minutes(arrival_time);
        return (arrival >= departure) ? (arrival - departure) : (arrival + 1440 - departure);
    }

    // 新增: 添加乘客到候补列表
    void add_to_waitlist(const Customer& customer) {
        waitlist.push_back(customer);
    }

    // 新增: 检查并处理候补列表
    bool process_waitlist() {
        if (!waitlist.empty() && remaining_tickets() > 0) {
            booked_customers.push_back(waitlist.front());
            waitlist.erase(waitlist.begin());
            return true;
        }
        return false;
    }
};

class FlightSystem {
    unordered_map<string, Flight> flights;
    unordered_map<string, vector<pair<string, int>>> graph;

public:
    void load_data(const string& filename) {
        ifstream infile(filename);
        if (!infile) {
            cerr << "无法打开文件: " << filename << endl;
            return;
        }

        string line;
        int line_number = 0;
        while (getline(infile, line)) {
            line_number++;
            istringstream iss(line);
            Flight flight;
            int customer_count;

            if (!(iss >> flight.flight_number >> flight.airplane_number
                >> flight.start >> flight.end >> flight.day_of_week
                >> flight.date >> flight.departure_time >> flight.arrival_time
                >> flight.crew_quota >> flight.total_tickets >> customer_count)) {
                cerr << "Error parsing line " << line_number << ": " << line << endl;
                cerr << "格式应为: 航班号 飞机号 起始城市 目的地 星期几 日期 出发时间 到达时间 乘员定额 总票量 已订票客户数 [客户名 订票数 舱位等级]..." << endl;
                continue;
            }

            for (int i = 0; i < customer_count; ++i) {
                Customer customer;
                if (!(iss >> customer.name >> customer.ticket_count >> customer.cabin_class)) {
                    cerr << "Error parsing customer info in line " << line_number << ": " << line << endl;
                    break;
                }
                flight.booked_customers.push_back(customer);
            }

            flights[flight.flight_number] = flight;
            graph[flight.start].push_back({ flight.end, flight.flight_duration() });
        }
    }
    void save_data(const string& filename) const {
        ofstream outfile(filename);
        if (!outfile) {
            cerr << "无法打开文件: " << filename << endl;
            return;
        }

        for (const auto& pair : flights) {
            const Flight& flight = pair.second;
            outfile << flight.flight_number << " "
                << flight.airplane_number << " "
                << flight.start << " "
                << flight.end << " "
                << flight.day_of_week << " "
                << flight.date << " "
                << flight.departure_time << " "
                << flight.arrival_time << " "
                << flight.crew_quota << " "
                << flight.total_tickets << " "
                << flight.booked_customers.size();

            for (const auto& customer : flight.booked_customers) {
                outfile << " " << customer.name << " "
                    << customer.ticket_count << " "
                    << customer.cabin_class;
            }

            outfile << endl;
        }
    }
    void add_flight(const Flight& flight) {
        flights[flight.flight_number] = flight;
        graph[flight.start].push_back({ flight.end, flight.flight_duration() });
    }
    // 按时间最短查询航线（考虑换乘等待时间）
    void find_shortest_time_route(const string& start, const string& end) const {
        unordered_map<string, int> min_time;
        unordered_map<string, string> previous;
        priority_queue<pair<int, string>, vector<pair<int, string>>, greater<>> pq;
        unordered_map<string, int> visit_count;

        for (const auto& node : graph) {
            min_time[node.first] = INT_MAX;
            visit_count[node.first] = 0;
        }
        if (graph.find(start) == graph.end()) {
            cout << "起点不存在于航线网络中。" << endl;
            return;
        }
        if (graph.find(end) == graph.end()) {
            cout << "终点不存在于航线网络中。" << endl;
            return;
        }
        min_time[start] = 0;
        pq.push({ 0, start });

        bool circular_route_found = false;
        vector<string> circular_route;
        int circular_route_time = 0;

        while (!pq.empty()) {
            int current_time = pq.top().first;
            string current_city = pq.top().second;
            pq.pop();

            visit_count[current_city]++;

            if (current_city == start && visit_count[current_city] > 1) {
                // 找到了从起点经过终点再回到起点的环路
                circular_route_found = true;
                circular_route = reconstruct_path(previous, start, end);
                circular_route.push_back(start);
                circular_route_time = current_time;
                break;
            }

            if (visit_count[current_city] > 2) continue; // 限制每个城市最多访问两次

            for (const auto& neighbor : graph.at(current_city)) {
                string next_city = neighbor.first;
                int travel_time = neighbor.second;

                int total_time = current_time + travel_time;

                if (total_time < min_time[next_city] ||
                    (next_city == start && visit_count[end] > 0)) { // 允许回到起点，但仅当已经访问过终点
                    min_time[next_city] = total_time;
                    previous[next_city] = current_city;
                    pq.push({ total_time, next_city });
                }
            }
        }

        // 输出最短时间路径
        if (min_time[end] == INT_MAX) {
            cout << "无法到达目的地。" << endl;
        }
        else {
            vector<string> shortest_path = reconstruct_path(previous, start, end);

            cout << "最短时间路径: " << min_time[end] << " 分钟" << endl;
            cout << "路径: ";
            for (const auto& city : shortest_path) {
                cout << city << " ";
            }
            cout << endl;
        }

        // 输出检测到的环路（如果存在）
        if (circular_route_found) {
            cout << "\n检测到从" << start << "经过" << end << "再回到" << start << "的环路：" << endl;
            cout << "环路: ";
            for (const auto& city : circular_route) {
                cout << city << " ";
            }
            cout << endl;
            cout << "环路总时间: " << circular_route_time << " 分钟" << endl;
        }
    }

    vector<string> reconstruct_path(const unordered_map<string, string>& previous,
        const string& start, const string& end) const {
        vector<string> path;
        for (string at = end; at != start; at = previous.at(at)) {
            path.push_back(at);
        }
        path.push_back(start);
        reverse(path.begin(), path.end());
        return path;
    }

    void query_routes(const string& start, const string& end, bool shortest_time) const {
        if (!shortest_time) {
            // 查询直达航线
            bool direct_found = false;
            for (const auto& pair : flights) {
                const Flight& flight = pair.second;
                if (flight.start == start && flight.end == end) {
                    cout << "直达航线 - 航班号: " << flight.flight_number
                        << ", 飞机号: " << flight.airplane_number
                        << ", 星期: " << flight.day_of_week
                        << ", 日期: " << flight.date
                        << ", 出发时间: " << flight.departure_time
                        << ", 到达时间: " << flight.arrival_time
                        << ", 飞行时间: " << flight.flight_duration() << " 分钟"
                        << ", 余票: " << flight.remaining_tickets() << endl;
                    direct_found = true;
                }
            }

            if (!direct_found) {
                cout << "没有直达航线，查找中转航线..." << endl;
            }

            // 使用修改后的 BFS 查找所有路径
            queue<vector<string>> paths;
            paths.push({ start });
            set<vector<string>> unique_paths;
            int path_limit = 10; // 限制输出的路径数量

            while (!paths.empty() && unique_paths.size() < path_limit) {
                auto current_path = paths.front();
                paths.pop();
                string last_city = current_path.back();

                if (last_city == end) {
                    if (unique_paths.find(current_path) == unique_paths.end()) {
                        unique_paths.insert(current_path);
                        print_route(current_path);
                    }
                    continue; // 找到一条路径后，不再继续扩展这条路径
                }

                if (current_path.size() > graph.size()) {
                    continue; // 如果路径长度超过城市总数，可能存在循环，停止扩展
                }

                if (graph.find(last_city) != graph.end()) {
                    for (const auto& neighbor : graph.at(last_city)) {
                        string next_city = neighbor.first;
                        if (find(current_path.begin(), current_path.end(), next_city) == current_path.end()) {
                            auto new_path = current_path;
                            new_path.push_back(next_city);
                            paths.push(new_path);
                        }
                    }
                }
            }

            if (unique_paths.empty()) {
                cout << "未找到可行的中转路径。" << endl;
            }
        }
        else {
            // 按时间最短查询
            find_shortest_time_route(start, end);
        }
    }

    void print_route(const vector<string>& route) const {
        cout << "路径: ";
        int total_duration = 0;
        for (size_t i = 0; i < route.size() - 1; ++i) {
            cout << route[i] << " -> ";
            for (const auto& pair : flights) {
                const Flight& flight = pair.second;
                if (flight.start == route[i] && flight.end == route[i + 1]) {
                    total_duration += flight.flight_duration();
                    break;
                }
            }
        }
        cout << route.back() << endl;
        cout << "中转次数: " << route.size() - 2 << ", 总飞行时间: " << total_duration << " 分钟" << endl << endl;
    }
    
    void book_ticket(const string& flight_number, const Customer& customer) {
        auto it = flights.find(flight_number);
        if (it != flights.end()) {
            Flight& flight = it->second;
            if (flight.remaining_tickets() >= customer.ticket_count) {
                flight.booked_customers.push_back(customer);
                cout << "订票成功。" << endl;
            }
            else {
                cout << "该航班已满。是否要加入候补列表？(Y/N): ";
                char choice;
                cin >> choice;
                if (toupper(choice) == 'Y') {
                    flight.add_to_waitlist(customer);
                    cout << "已添加到候补列表。" << endl;
                }
                else {
                    cout << "订票失败。" << endl;
                }
            }
        }
        else {
            cout << "航班号不存在。" << endl;
        }
    }

    void refund_ticket(const string& flight_number, const string& customer_name, const string& date) {
        auto it = flights.find(flight_number);
        if (it != flights.end()) {
            Flight& flight = it->second;
            if (flight.date == date) {
                auto& customers = flight.booked_customers;
                auto cust_it = find_if(customers.begin(), customers.end(),
                    [&](const Customer& c) { return c.name == customer_name; });
                if (cust_it != customers.end()) {
                    customers.erase(cust_it);
                    cout << "退票成功。" << endl;

                    // 处理候补列表
                    if (flight.process_waitlist()) {
                        cout << "候补乘客已自动订票。" << endl;
                    }
                }
                else {
                    cout << "未找到该客户的订票记录。" << endl;
                }
            }
            else {
                cout << "日期不匹配。" << endl;
            }
        }
        else {
            cout << "航班号不存在。" << endl;
        }
    }
    
    void display_waitlist(const string& flight_number) const {
        auto it = flights.find(flight_number);
        if (it != flights.end()) {
            const Flight& flight = it->second;
            cout << "航班 " << flight_number << " 的候补列表：" << endl;
            for (const auto& customer : flight.waitlist) {
                cout << "姓名: " << customer.name
                    << ", 票数: " << customer.ticket_count
                    << ", 舱位等级: " << customer.cabin_class << endl;
            }
        }
        else {
            cout << "航班号不存在。" << endl;
        }
    }
    void display_all_flights() const {
        for (const auto& pair : flights) {
            const Flight& flight = pair.second;
            try {
                int booked_tickets = 0;
                for (const auto& customer : flight.booked_customers) {
                    booked_tickets += customer.ticket_count;
                }
                int available_seats = flight.total_tickets - flight.crew_quota;
                int remaining_tickets = available_seats - booked_tickets;

                cout << "航班号: " << flight.flight_number
                    << ", 起始城市: " << flight.start
                    << ", 目的地: " << flight.end
                    << ", 飞机号: " << flight.airplane_number
                    << ", 星期: " << flight.day_of_week
                    << ", 日期: " << flight.date
                    << ", 出发时间: " << flight.departure_time
                    << ", 到达时间: " << flight.arrival_time
                    << ", 飞行时间: " << flight.flight_duration() << " 分钟"
                    << ", 乘员定额: " << flight.crew_quota
                    << ", 总票量: " << flight.total_tickets
                    << ", 可售票数: " << available_seats
                    << ", 已订票数: " << booked_tickets
                    << ", 余票: " << remaining_tickets << endl;
            }
            catch (const std::exception& e) {
                cerr << "处理航班 " << flight.flight_number << " 时发生错误: " << e.what() << endl;
            }
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
    cout << "6. 查看候补列表" << endl;  // 新选项
    cout << "退出请输入Q" << endl;
    cout << "请选择操作 (1-6): ";
}

int main() {
    FlightSystem system;
    system.load_data("flights_data.txt");

    char choice;
    while (true) {
        display_menu();
        cin >> choice;
        cin.ignore(); // 清除输入缓冲区

        switch (choice) {
        case '1': {
            string start, end;
            char query_type;
            cout << "请输入起点城市代码: ";
            cin >> start;
            cout << "请输入目的地代码: ";
            cin >> end;
            cout << "请选择查询方式：\n"
                << "A. 查询直达航线\n"
                << "B. 查询按时间最短的航线\n"
                << "请输入 (A/B): ";
            cin >> query_type;

            bool shortest_time = (query_type == 'B' || query_type == 'b');
            system.query_routes(start, end, shortest_time);
            break;
        }

        case '2': {
            string flight_number, customer_name,date;
            int ticket_count, cabin_class;
            cout << "请输入航班号: ";
            cin >> flight_number;
            cout << "请输入客户姓名: ";
            cin >> customer_name;
            cout << "请输入订票数量: ";
            cin >> ticket_count;
            cout << "请输入舱位等级（1, 2 或 3）: ";
            cin >> cabin_class;

            // 验证输入的票数和舱位等级
            if (ticket_count <= 0 || cabin_class < 1 || cabin_class > 3) {
                cout << "无效的票数或舱位等级。" << endl;
                break;
            }

            Customer customer = { customer_name, ticket_count, cabin_class };
            system.book_ticket(flight_number, customer);
            break;
        }
        case '3': {
            string flight_number, customer_name, date;
            cout << "请输入航班号: ";
            cin >> flight_number;
            cout << "请输入客户姓名: ";
            cin >> customer_name;
            cout << "请输入具体日期: ";
            cin >> date;
            system.refund_ticket(flight_number, customer_name, date);
            break;
        }
        case '4': {
            Flight flight;
            cout << "请输入始发地: ";
            cin >> flight.start;
            cout << "请输入目的地: ";
            cin >> flight.end;
            cout << "请输入航班号: ";
            cin >> flight.flight_number;
            cout << "请输入飞机号: ";
            cin >> flight.airplane_number;
            cout << "请输入星期几（如 Monday）: ";
            cin >> flight.day_of_week;
            cout << "请输入年月日: ";
            cin >> flight.date;
            cout << "请输入出发时间 (HH:MM): ";
            cin >> flight.departure_time;
            cout << "请输入到达时间 (HH:MM): ";
            cin >> flight.arrival_time;
            cout << "请输入乘员定额: ";
            cin >> flight.crew_quota;
            cout << "请输入总票量: ";
            cin >> flight.total_tickets;
            flight.booked_customers.clear(); // 新航班没有已订票客户

            system.add_flight(flight);
            break;
        }
        case '5': {
            system.display_all_flights();
            break;
        }
        case '6': {
            string flight_number;
            cout << "请输入航班号: ";
            cin >> flight_number;
            system.display_waitlist(flight_number);
            break;
        }
        case 'Q': {
            system.save_data("flights_data.txt");
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