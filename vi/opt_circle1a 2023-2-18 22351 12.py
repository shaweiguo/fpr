#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
Example of 1D fitting to data: Circle radius

@author: tim
"""

from scipy.optimize import minimize
import numpy as np
import matplotlib.pyplot as plt

def make_circle1(n, r):
  """Return n points on circle centred at (0,0) with radius r"""
  
  t=np.linspace(0,2*np.pi,n,endpoint=False)
  p=np.zeros((2,n))
  p[0,:]=r*np.cos(t)
  p[1,:]=r*np.sin(t)
  return p


def add_noise(p, sd):
  """Add gaussian noise with standard deviation of sd to elements of p"""
  
  return p+np.random.normal(size=p.shape,scale=sd)


# Create points for circle (radius 2.5)
pts=make_circle1(50, 2.5)

plt.plot(pts[0],pts[1])

pts2=add_noise(pts,0.05)
plt.plot(pts2[0],pts2[1],"+")

