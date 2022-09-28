# Go-image

## Introduction

This go package has been created to do image processing with golang. I'm currently working to improve it (parallels proccesses, etc). In all exemple, the default image is :

![default](./assets/images/gopher_and_docker.jpeg)

## Table of Content

1. [Effects](#effects)
    - [Grayscale](#grayscale)
    - [Threshold](#threshold)
    - [Invert](#invert)
    - [Sobel Edge detection](#sobal-edge-detection)
    - [Gaussian Blur](#gaussian-blur)
2. [Licence](#licence)

## Effects

### Grayscale

You can transform a picture in grayscale with the Grayscale function.

```golang
newImg := effects.Grayscale(img)
```

![exemple](./assets/images/Grayscale.jpeg)

By default the Grayscale transformation uses the ITU-R recomandation but you can modify it. For exemple :

```golang
newImg := effects.Grayscale(img, 0.2, 0.5, 0.3)
```

![exemple](./assets/images/configGrayscale.jpeg)

Please when you use your own config, make sure to use all three parameters and make sure that the sum is equal to one.

### Threshold

You can threshold a picture. It means you fix a gray level and every pixel under that gray level will be black. Others will be white.

```golang
newImg := effects.Threshold(img, 0) // It will render a fully white image
```

![exemple](./assets/images/whiteThreshold.jpeg)

Here is an exemple of a half Threshold :

```golang
newImg := effects.Threshold(img, 187) // It will render a half Threshold image
```

![exemple](./assets/images/halfThreshold.jpeg)

### Invert

You can invert an RGBA image.

```golang
newImg := effects.Invert(img)
```

![exemple](./assets/images/Invert.jpeg)

You can also select a square to invert inside the image with the config array :

```golang
newImg := effects.Invert(img, 0, 200, 0, 200) // It will invert top left square of 3x3 pixels if the image size is large enough
```

![exemple](./assets/images/partialInvert.jpeg)

### Sobal Edge detection

You can use the Sobal edge detection method.

```golang
newImg := effects.SobalEdge(img)
```

![exemple](./assets/images/SobalEdge.jpeg)

### Gaussian Blur

You can blur an image. You can choose the blur radius. The bigger the blur radius, the more blur the picture will be.

```golang
newImg := effects.GaussianBlur(img, 10)
```

![exemple](./assets/images/Blur.jpeg)

## Licence

This project is licensed under the MIT license. Please read the LICENSE file.
