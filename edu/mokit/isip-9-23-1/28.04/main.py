a = 10
print(float(a))
b = "0"
try:
    print(a + b)
except TypeError:
    pass
except ZeroDivisionError:
    pass
except:
    pass
