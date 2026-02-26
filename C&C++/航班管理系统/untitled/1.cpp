#include <iostream>
#include <vector>
#include <string>
#include <fstream>
#include <sstream>
#include <json/json.h> // 使用 jsoncpp 库处理 JSON 数据

using namespace std;

class Ticket {
public:
    string customerName;
    int quantity;
    int classLevel;
    vector<string> seatNumbers;

    Ticket(string name, int qty, int level, vector<string> seats)
            : customerName(name), quantity(qty), classLevel(level), seatNumbers(seats) {}
};

class Flight {
public:
    string flightNumber;
    string destination;
    string dayOfWeek;
    int seatCapacity;
    vector<Ticket> bookedTickets;

    Flight(string number, string dest, string day, int capacity)
            : flightNumber(number), destination(dest), dayOfWeek(day), seatCapacity(capacity) {}

    int availableSeats() {
        int booked = 0;
        for (const auto& ticket : bookedTickets) {
            booked += ticket.quantity;
        }
        return seatCapacity - booked;
    }

    vector<string> bookTicket(string customerName, int quantity, int classLevel) {
        if (availableSeats() >= quantity) {
            vector<string> seatNumbers;
            for (int i = 1; i <= quantity; ++i) {
                seatNumbers.push_back(flightNumber + "-" + to_string(i));
            }
            bookedTickets.emplace_back(customerName, quantity, classLevel, seatNumbers);
            return seatNumbers;
        }
        return {};
    }

    bool cancelTicket(string customerName) {
        for (auto it = bookedTickets.begin(); it != bookedTickets.end(); ++it) {
            if (it->customerName == customerName) {
                bookedTickets.erase(it);
                return true;
            }
        }
        return false;
    }
};

class FlightSystem {
private:
    vector<Flight> flights;

public:
    void loadData(const string& filename) {
        ifstream file(filename);
        Json::Value data;
        file >> data;

        for (const auto& flight : data["flights"]) {
            flights.emplace_back(flight["flight_number"].asString(),
                                 flight["destination"].asString(),
                                 flight["day_of_week"].asString(),
                                 flight["seat_capacity"].asInt());
        }
    }

    void queryFlight(const string& destination) {
        for (const auto& flight : flights) {
            if (flight.destination == destination) {
                cout << "航班号: " << flight.flightNumber << ", 星期: " << flight.dayOfWeek
                     << ", 余票: " << flight.availableSeats() << endl;
            }
        }
    }

    vector<string> bookTicket(const string& flightNumber, const string& customerName, int quantity, int classLevel) {
        for (auto& flight : flights) {
            if (flight.flightNumber == flightNumber) {
                return flight.bookTicket(customerName, quantity, classLevel);
            }
        }
        return {};
    }

    bool cancelTicket(const string& customerName, const string& flightNumber) {
        for (auto& flight : flights) {
            if (flight.flightNumber == flightNumber) {
                return flight.cancelTicket(customerName);
            }
        }
        return false;
    }

    void displayMenu() {
        cout << "1. 查询航线" << endl;
        cout << "2. 订票" << endl;
        cout << "3. 退票" << endl;
        cout << "4. 退出" << endl;
    }

    void run() {
        while (true) {
            displayMenu();
            int choice;
            cout << "请选择操作: ";
            cin >> choice;

            if (choice == 1) {
                string destination;
                cout << "请输入终点站名: ";
                cin >> destination;
                queryFlight(destination);
            } else if (choice == 2) {
                string flightNumber, customerName;
                int quantity, classLevel;
                cout << "请输入航班号: ";
                cin >> flightNumber;
                cout << "请输入客户姓名: ";
                cin >> customerName;
                cout << "请输入订票数量: ";
                cin >> quantity;
                cout << "请输入舱位等级 (1, 2, 3): ";
                cin >> classLevel;

                vector<string> seats = bookTicket(flightNumber, customerName, quantity, classLevel);
                if (!seats.empty()) {
                    cout << "订票成功，座位号: ";
                    for (const auto& seat : seats) {
                        cout << seat << " ";
                    }
                    cout << endl;
                } else {
                    cout << "订票失败，余票不足。" << endl;
                }
            } else if (choice == 3) {
                string customerName, flightNumber;
                cout << "请输入客户姓名: ";
                cin >> customerName;
                cout << "请输入航班号: ";
                cin >> flightNumber;

                if (cancelTicket(customerName, flightNumber)) {
                    cout << "退票成功。" << endl;
                } else {
                    cout << "退票失败。" << endl;
                }
            } else if (choice == 4) {
                cout << "退出系统。" << endl;
                break;
            } else {
                cout << "无效选择，请重试。" << endl;
            }
        }
    }
};

int main() {
    FlightSystem system;
    system.loadData("flights.json"); // 假设数据存储在 flights.json 文件中
    system.run();
    return 0;
}