#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
Example of optimizing an analytic function

@author: tim
"""

from scipy.optimize import brent
import matplotlib.pyplot as plt
import numpy as np

# Define a quadratic function
def f1(x):
    return x*x + 2*x - 3

# Plot it
t=np.linspace(-3,3,50)
y=f1(t)

plt.plot(t,f1(t))
plt.show()

# Find x for minimum
x_min=brent(f1, brack=(-3,3))
print("x_min: ",x_min,"  f(x_min)=", f1(x_min))