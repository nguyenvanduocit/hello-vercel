# hello-vercel

Play around with Vercel Serverless Function

There are two functions:

1. favicon
1. image/random

## favicon

The function to get favicon of a url.

### Usage

```
curl -L "https://hello-vercel-gold.vercel.app/api/favicon?url=https://12bit.vn" --output ./icon.png
```

## image/random

The function to get a random image.

### Usage

```
curl -L "https://hello-vercel-gold.vercel.app/api/image/random" --output ./random.jpg
```
