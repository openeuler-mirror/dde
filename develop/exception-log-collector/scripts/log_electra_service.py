#!/usr/bin/python3
# -*- coding: utf-8 -*-

# scripts/log_electra_service.py

from flask import Flask, request, jsonify
from transformers import AutoTokenizer, AutoModelForSequenceClassification
import torch

app = Flask(__name__)

# 加载模型
model_name = "sustcsonglin/LogELECTRA"
tokenizer = AutoTokenizer.from_pretrained(model_name)
model = AutoModelForSequenceClassification.from_pretrained(model_name)


@app.route("/analyze", methods=["POST"])
def analyze():
    data = request.get_json()
    log_line = data.get("log", "")
    inputs = tokenizer(log_line, return_tensors="pt", truncation=True)
    outputs = model(**inputs)
    logits = outputs.logits
    probabilities = torch.softmax(logits, dim=1)
    predicted_label = torch.argmax(probabilities, dim=1).item()
    confidence = probabilities[0][predicted_label].item()
    result = {"label": predicted_label, "confidence": confidence}
    return jsonify(result)


if __name__ == "__main__":
    app.run(host="0.0.0.0", port=5000)
