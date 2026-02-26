#include "mainwindow.h"
#include "ui_mainwindow.h"
#include <QMessageBox>
#include <QFile>
#include <QTextStream>

MainWindow::MainWindow(QWidget *parent)
    : QMainWindow(parent)
    , ui(new Ui::MainWindow) {
    ui->setupUi(this);
    loadUsers();
}

MainWindow::~MainWindow() {
    delete ui;
}

void MainWindow::on_registerButton_clicked() {
    QString username = ui->usernameLineEdit->text();
    QString phoneNumber = ui->phoneLineEdit->text();
    QString password = ui->passwordLineEdit->text();

    if (username.isEmpty() || phoneNumber.isEmpty() || password.isEmpty()) {
        QMessageBox::warning(this, "Registration Failed", "All fields are required.");
        return;
    }

    for (const User &user : users) {
        if (user.username == username) {
            QMessageBox::warning(this, "Registration Failed", "Username already exists.");
            return;
        }
    }

    User newUser = {username, phoneNumber, password};
    users.append(newUser);
    saveUsers();

    QMessageBox::information(this, "Registration Successful", "You have successfully registered.");
}

void MainWindow::on_loginButton_clicked() {
    QString username = ui->usernameLineEdit->text();
    QString password = ui->passwordLineEdit->text();

    for (const User &user : users) {
        if (user.username == username && user.password == password) {
            currentUser = user;
            QMessageBox::information(this, "Login Successful", "You have successfully logged in.");
            return;
        }
    }

    QMessageBox::warning(this, "Login Failed", "Invalid username or password.");
}

void MainWindow::on_bookSeatButton_clicked() {
    if (currentUser.username.isEmpty()) {
        QMessageBox::warning(this, "Booking Failed", "Please login first.");
        return;
    }

    QString reservationTime = ui->reservationTimeLineEdit->text();
    QString seatType = ui->seatTypeLineEdit->text();

    if (reservationTime.isEmpty() || seatType.isEmpty()) {
        QMessageBox::warning(this, "Booking Failed", "All fields are required.");
        return;
    }

    QMessageBox::information(this, "Booking Successful", "Seat booked successfully.");
}

void MainWindow::on_cancelBookingButton_clicked() {
    if (currentUser.username.isEmpty()) {
        QMessageBox::warning(this, "Cancellation Failed", "Please login first.");
        return;
    }

    QMessageBox::information(this, "Cancellation Successful", "Booking cancelled successfully.");
}

void MainWindow::on_modifyBookingButton_clicked() {
    if (currentUser.username.isEmpty()) {
        QMessageBox::warning(this, "Modification Failed", "Please login first.");
        return;
    }

    QMessageBox::information(this, "Modification Successful", "Booking modified successfully.");
}

void MainWindow::loadUsers() {
    QFile file("users.txt");
    if (file.open(QIODevice::ReadOnly | QIODevice::Text)) {
        QTextStream in(&file);
        while (!in.atEnd()) {
            QStringList fields = in.readLine().split(" ");
            if (fields.size() == 3) {
                User user = {fields[0], fields[1], fields[2]};
                users.append(user);
            }
        }
        file.close();
    }
}

void MainWindow::saveUsers() {
    QFile file("users.txt");
    if (file.open(QIODevice::WriteOnly | QIODevice::Text)) {
        QTextStream out(&file);
        for (const User &user : users) {
            out << user.username << " " << user.phoneNumber << " " << user.password << "\n";
        }
        file.close();
    }
}
