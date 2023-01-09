import tensorflow as tf
from tensorflow.keras.layers import Input

x = Input(batch_shape=(100, 100))
print(x)
