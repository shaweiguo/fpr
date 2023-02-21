#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
Example of 1D fitting to data: Circle radius

@author: tim
"""

from scipy.optimize import brent
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

class CostFunction_circle1:
  """Cost function for circle fit (x=[r]). Initialised with points."""
  
  def __init__(self,pts):
    self.pts=pts
  
  def f(self,x):
      """Evaluate cost function fitting circle radius x"""
      r2=np.square(self.pts[0,:]) + np.square(self.pts[1,:])
      d=np.square(np.sqrt(r2)-x)
      return np.sum(d)

# Create points for circle (radius 2.5)
n_pts=40
true_r=3.5
pts=make_circle1(n_pts, true_r)

plt.plot(pts[0],pts[1])

pts2=add_noise(pts,0.25)
plt.plot(pts2[0],pts2[1],"+")
# Arrange that axis steps are same, so circle not squashed
plt.axis('equal')
plt.show()

# Set up cost function for circle fit
c1=CostFunction_circle1(pts2)
r1=brent(c1.f,brack=(0,10))
print("True radius={:0.2f}  Estimated radius={:0.2f}".format(true_r,r1))

