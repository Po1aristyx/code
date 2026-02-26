#include <iostream>
#include <iomanip>
#include <vector>
#include <string>
using namespace std;

// 基类 User，表示公司员工的基本信息
class User {
protected:
    string name;
    string id;
    int age;
    bool sex; // true 表示男，false 表示女
    float salary;
public:
    User(string& a, string& b, int c, bool d) :name(a), id(b), age(c), sex(d) {}
    virtual void getpay() = 0; // 纯虚函数，计算薪金
    virtual void print() = 0;  // 纯虚函数，打印员工信息
    void basicprint() {
        cout << setw(6) << "职工编号:" << setw(6) << id << " 姓名：" << setw(6) << name << setw(4) << "  月薪：" << setw(6) << salary << "元" << endl;
    }
};

// 派生类 Boss，表示老板
class Boss :public User {
private:
    int annual_salary;
public:
    void getpay() override {
        salary = annual_salary / 12.0f; // 计算月薪
    }

    void print() override {
        cout << setw(6) << "职工编号:" << setw(6) << id << setw(4) << " 职务：老板" << setw(4) << " 姓名：" << setw(6) << name << setw(4) << " 年龄：" << setw(4) << age << setw(5) << " 性别：" << setw(4) << (sex ? "男" : "女") << setw(5) << " 月薪：" << setw(6) << salary << "元" << endl;
    }

    Boss(string& a, string& b, int c, bool d, int e) : User(a, b, c, d), annual_salary(e) { getpay(); }
};

// 派生类 Employee，表示雇员
class Employee :public User {
private:
    int basic_salary;
    int bonus;
public:
    void getpay() override {
        salary = basic_salary + bonus; // 计算月薪
    }

    void print() override {
        cout << setw(6) << "职工编号:" << setw(6) << id << setw(4) << " 职务：雇员" << setw(4) << " 姓名：" << setw(6) << name << setw(4) << " 年龄：" << setw(4) << age << setw(5) << " 性别：" << setw(4) << (sex ? "男" : "女") << setw(5) << " 月薪：" << setw(6) << salary << "元" << endl;
    }

    Employee(string& a, string& b, int c, bool d, int e, int f) : User(a, b, c, d), basic_salary(e), bonus(f) { getpay(); }
};

// 派生类 Hourly_Worker，表示小时工
class Hourly_Worker :public User {
private:
    int time;
    int time_salary;
public:
    void getpay() override {
        salary = time * time_salary; // 计算工资
    }

    void print() override {
        cout << setw(6) << "职工编号:" << setw(6) << id << setw(4) << " 职务：小时工" << setw(4) << " 姓名：" << setw(6) << name << setw(4) << " 年龄：" << setw(4) << age << setw(5) << " 性别：" << setw(4) << (sex ? "男" : "女") << setw(5) << " 月薪：" << setw(6) << salary << "元" << endl;
    }

    Hourly_Worker(string& a, string& b, int c, bool d, int e, int f) :User(a, b, c, d), time(e), time_salary(f) { getpay(); }
};

// 派生类 CommWorker，表示营销人员
class CommWorker :public User {
private:
    int basic_salary;
    float profit;
public:
    void getpay() override {
        salary = basic_salary + 0.05 * profit; // 计算工资
    }

    void print() override {
        cout << setw(6) << "职工编号:" << setw(6) << id << setw(4) << " 职务：营销人员" << setw(4) << " 姓名：" << setw(6) << name << setw(4) << " 年龄：" << setw(4) << age << setw(5) << " 性别：" << setw(4) << (sex ? "男" : "女") << setw(5) << " 月薪：" << setw(6) << salary << "元" << endl;
    }

    CommWorker(string& a, string& b, int c, bool d, int e, float f) :User(a, b, c, d), basic_salary(e), profit(f) { getpay(); }
};

// 报表类 Report，管理和输出员工信息
class Report {
private:
    vector<User*> list; // 员工列表
public:
    bool add() {
        cout << "请输入员工类型！(1-老板 2-雇员 3-小时工 4-营销人员)或者输入0退出添加程序！" << endl;
        int lx, age, a, b;
        string name, id;
        bool sex;
        float c;
        cin >> lx;
        if (lx > 4 || lx < 0) {
            cout << "输入类型错误！请重新输入！" << endl;
            return true;
        }
        if (lx == 0)
            return false;

        cout << "请输入员工姓名!" << endl;
        cin >> name;
        cout << "请输入员工编号！" << endl;
        cin >> id;
        cout << "请输入员工年龄！" << endl;
        cin >> age;
        if (age < 14 || age > 60) {
            cout << "该员工年龄非法！拒绝录用！" << endl;
            return true;
        }
        cout << "请输入员工性别！(0-女，1-男)" << endl;
        cin >> sex;

        switch (lx) {
        case 1:
            cout << "请输入年薪！" << endl;
            cin >> a;
            list.push_back(new Boss(name, id, age, sex, a));
            break;
        case 2:
            cout << "请输入基本工资！" << endl;
            cin >> a;
            cout << "请输入奖金！" << endl;
            cin >> b;
            list.push_back(new Employee(name, id, age, sex, a, b));
            break;
        case 3:
            cout << "请输入工作时间！" << endl;
            cin >> a;
            cout << "请输入时薪！" << endl;
            cin >> b;
            list.push_back(new Hourly_Worker(name, id, age, sex, a, b));
            break;
        case 4:
            cout << "请输入基本工资！" << endl;
            cin >> a;
            cout << "请输入销售利润！" << endl;
            cin >> c;
            list.push_back(new CommWorker(name, id, age, sex, a, c));
            break;
        default:
            break;
        }
        return true;
    }

    void print() {
        cout << "公司每个人员当月薪金情况报表:" << endl;
        for (int i = 0; i < list.size(); i++) {
            list[i]->basicprint();
        }
    }
};

int main() {
    Report re;
    int ml;
    while (true) {
        system("cls");
        cout << "----------------------" << endl;
        cout << "-----职工管理系统-----" << endl;
        cout << "-----1  添加员工 -----" << endl;
        cout << "-----2  输出报表 -----" << endl;
        cout << "-----0  退出系统 -----" << endl;
        cout << "----------------------" << endl;
        cout << "请输入操作！" << endl;
        cin >> ml;
        switch (ml) {
        case 1:
            while (re.add());
            break;
        case 2:
            re.print();
            break;
        case 0:
            exit(0);
            break;
        default:
            break;
        }
        system("pause");
    }
    return 0;
}
