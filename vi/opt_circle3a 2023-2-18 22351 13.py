#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
Example of fitting a circle with arbitrary centre and radius to data.

@author: tim
"""

from scipy.optimize import minimize
import numpy as np
import matplotlib.pyplot as plt


def make_circle3(n, x,y,r):
  """Return n points on circle centred at (x,y) with radius r"""
  
  t=np.linspace(0,2*np.pi,n,endpoint=False)
  p=np.zeros((2,n))
  p[0,:]=x+r*np.cos(t)
  p[1,:]=y+r*np.sin(t)
  return p


def add_noise(p, sd):
  """Add gaussian noise with standard deviation of sd to elements of p"""
  
  return p+np.random.normal(size=p.shape,scale=sd)

class CostFunction_circle3:
  """Cost function for circle fit (x=[cx,cy,r]. Initialised with points."""
  
  def __init__(self,pts):
    self.pts=pts
  
  def f(self,x):
      """Evaluate cost function fitting circle centre (x[0],x[1]) radius x[2]"""
      r2=np.square(self.pts[0,:]-x[0]) + np.square(self.pts[1,:]-x[1])
      d=np.square(np.sqrt(r2)-x[2])
      return np.sum(d)

# Generate points on a circle centre (3,4) radius 5
#Fitting a function of 3 parameters
pts=make_circle3(25, 3,4,5)
pts2=add_noise(pts,0.1)

plt.plot(pts2[0],pts2[1],"+")
plt.show()

# Set up a cost function to evaluate how well a circle fits to pts2
c3=CostFunction_circle3(pts2)

# Initialise a 3 element vector to zero (as a start)
x0=np.zeros(3)

# Use Powell's method to fit, starting at x
res=mi(c3.f,x0,method='Powell') 

print("Best fit has centre (",res.x[0],",",res.x[1],") radius ",res.x[2])

