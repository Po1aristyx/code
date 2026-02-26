import sys
print(sys.path)
try:
    import mysql.connector
    print("mysql.connector is imported successfully.")
except ModuleNotFoundError:
    print("mysql.connector is not found.")