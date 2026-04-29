a = 2
b = "0"
try:
    print(a / b)
except ZeroDivisionError:
    pass
except TypeError:
    pass
except:
    pass

print("end")
