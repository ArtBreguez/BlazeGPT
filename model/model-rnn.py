import json
import numpy as np
import tensorflow as tf
from tensorflow.keras.models import Sequential
from tensorflow.keras.layers import LSTM, Dense

# Carrega os dados do arquivo JSONL
X = []
y = []
with open("data.jsonl", "r") as f:
    for line in f:
        data = json.loads(line)
        X.append([int(x) for x in data["input"].split(',')])
        y.append(int(data["output"]))

# Converte as entradas para arrays numpy
X = np.array(X)

# Codifica as saídas como one-hot vectors
y = tf.keras.utils.to_categorical(y)

# Define o tamanho dos dados de entrada e saída
timesteps = X.shape[1]
features = 1
classes = y.shape[1]

# Define a arquitetura do modelo
model = Sequential([
    LSTM(64, input_shape=(timesteps, features), return_sequences=True),
    LSTM(64),
    Dense(classes, activation='softmax')
])

# Compila o modelo
model.compile(optimizer='adam', loss='categorical_crossentropy', metrics=['accuracy'])

# Treina o modelo
model.fit(X, y, epochs=10, batch_size=32)
