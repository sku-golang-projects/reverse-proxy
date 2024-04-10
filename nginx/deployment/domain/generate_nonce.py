#!/usr/bin/env python3
import base64
import os

def generate_nonce():
    random_bytes = os.urandom(16)
    nonce = base64.b64encode(random_bytes).decode('utf-8')
    nonce = nonce.replace('+', '').replace('=', '').replace('/', '')
    return nonce

if __name__ == "__main__":
    nonce = generate_nonce()
    print("Generated nonce:", nonce)