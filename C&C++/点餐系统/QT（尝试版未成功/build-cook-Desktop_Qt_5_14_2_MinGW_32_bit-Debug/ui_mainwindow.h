/********************************************************************************
** Form generated from reading UI file 'mainwindow.ui'
**
** Created by: Qt User Interface Compiler version 5.14.2
**
** WARNING! All changes made in this file will be lost when recompiling UI file!
********************************************************************************/

#ifndef UI_MAINWINDOW_H
#define UI_MAINWINDOW_H

#include <QtCore/QVariant>
#include <QtWidgets/QApplication>
#include <QtWidgets/QLineEdit>
#include <QtWidgets/QMainWindow>
#include <QtWidgets/QMenuBar>
#include <QtWidgets/QPushButton>
#include <QtWidgets/QStatusBar>
#include <QtWidgets/QWidget>

QT_BEGIN_NAMESPACE

class Ui_MainWindow
{
public:
    QWidget *centralwidget;
    QLineEdit *usernameLineEdit;
    QLineEdit *phoneLineEdit;
    QLineEdit *passwordLineEdit;
    QLineEdit *reservationTimeLineEdit;
    QLineEdit *seatTypeLineEdit;
    QPushButton *registerButton;
    QPushButton *loginButton;
    QPushButton *bookSeatButton;
    QPushButton *cancelBookingButton;
    QPushButton *modifyBookingButton;
    QMenuBar *menubar;
    QStatusBar *statusbar;

    void setupUi(QMainWindow *MainWindow)
    {
        if (MainWindow->objectName().isEmpty())
            MainWindow->setObjectName(QString::fromUtf8("MainWindow"));
        MainWindow->resize(800, 600);
        centralwidget = new QWidget(MainWindow);
        centralwidget->setObjectName(QString::fromUtf8("centralwidget"));
        usernameLineEdit = new QLineEdit(centralwidget);
        usernameLineEdit->setObjectName(QString::fromUtf8("usernameLineEdit"));
        usernameLineEdit->setGeometry(QRect(410, 100, 113, 21));
        phoneLineEdit = new QLineEdit(centralwidget);
        phoneLineEdit->setObjectName(QString::fromUtf8("phoneLineEdit"));
        phoneLineEdit->setGeometry(QRect(410, 130, 113, 21));
        passwordLineEdit = new QLineEdit(centralwidget);
        passwordLineEdit->setObjectName(QString::fromUtf8("passwordLineEdit"));
        passwordLineEdit->setGeometry(QRect(410, 160, 113, 21));
        reservationTimeLineEdit = new QLineEdit(centralwidget);
        reservationTimeLineEdit->setObjectName(QString::fromUtf8("reservationTimeLineEdit"));
        reservationTimeLineEdit->setGeometry(QRect(410, 190, 113, 21));
        seatTypeLineEdit = new QLineEdit(centralwidget);
        seatTypeLineEdit->setObjectName(QString::fromUtf8("seatTypeLineEdit"));
        seatTypeLineEdit->setGeometry(QRect(410, 220, 113, 21));
        registerButton = new QPushButton(centralwidget);
        registerButton->setObjectName(QString::fromUtf8("registerButton"));
        registerButton->setGeometry(QRect(110, 70, 93, 28));
        loginButton = new QPushButton(centralwidget);
        loginButton->setObjectName(QString::fromUtf8("loginButton"));
        loginButton->setGeometry(QRect(90, 120, 93, 28));
        bookSeatButton = new QPushButton(centralwidget);
        bookSeatButton->setObjectName(QString::fromUtf8("bookSeatButton"));
        bookSeatButton->setGeometry(QRect(100, 160, 93, 28));
        cancelBookingButton = new QPushButton(centralwidget);
        cancelBookingButton->setObjectName(QString::fromUtf8("cancelBookingButton"));
        cancelBookingButton->setGeometry(QRect(80, 190, 93, 28));
        modifyBookingButton = new QPushButton(centralwidget);
        modifyBookingButton->setObjectName(QString::fromUtf8("modifyBookingButton"));
        modifyBookingButton->setGeometry(QRect(90, 230, 93, 28));
        MainWindow->setCentralWidget(centralwidget);
        menubar = new QMenuBar(MainWindow);
        menubar->setObjectName(QString::fromUtf8("menubar"));
        menubar->setGeometry(QRect(0, 0, 800, 26));
        MainWindow->setMenuBar(menubar);
        statusbar = new QStatusBar(MainWindow);
        statusbar->setObjectName(QString::fromUtf8("statusbar"));
        MainWindow->setStatusBar(statusbar);

        retranslateUi(MainWindow);

        QMetaObject::connectSlotsByName(MainWindow);
    } // setupUi

    void retranslateUi(QMainWindow *MainWindow)
    {
        MainWindow->setWindowTitle(QCoreApplication::translate("MainWindow", "MainWindow", nullptr));
        registerButton->setText(QCoreApplication::translate("MainWindow", "PushButton", nullptr));
        loginButton->setText(QCoreApplication::translate("MainWindow", "PushButton", nullptr));
        bookSeatButton->setText(QCoreApplication::translate("MainWindow", "PushButton", nullptr));
        cancelBookingButton->setText(QCoreApplication::translate("MainWindow", "PushButton", nullptr));
        modifyBookingButton->setText(QCoreApplication::translate("MainWindow", "PushButton", nullptr));
    } // retranslateUi

};

namespace Ui {
    class MainWindow: public Ui_MainWindow {};
} // namespace Ui

QT_END_NAMESPACE

#endif // UI_MAINWINDOW_H
