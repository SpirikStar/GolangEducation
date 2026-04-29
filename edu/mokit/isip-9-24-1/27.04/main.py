a = 1
b = "0j"

try:
    print(a + b)
except TypeError:
    b = int(b)
except ZeroDivisionError:
    pass
except:
    pass
print("end")
