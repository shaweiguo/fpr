# This is a sample Python script.

# Press Shift+F10 to execute it or replace it with your code.
# Press Double Shift to search everywhere for classes, files, tool windows, actions, and settings.
import cv2 as cv


def show(name):
    # Use a breakpoint in the code line below to debug your script.
    # print(f'Hi, {name}')  # Press Ctrl+F8 to toggle the breakpoint.
    img = cv.imread('data/' + name + '.jpg')
    cv.imshow(name, img)
    cv.waitKey(0)


def show_grayscale(name):
    # Use a breakpoint in the code line below to debug your script.
    # print(f'Hi, {name}')  # Press Ctrl+F8 to toggle the breakpoint.
    img = cv.imread('data/' + name + '.jpg', cv.IMREAD_GRAYSCALE)
    cv.imshow(name, img)
    cv.waitKey(0)


def show_butterfly():
    capture = cv.VideoCapture('data/butterfly.mp4')
    isTrue = True
    while isTrue:
        isTrue, frame = capture.read()
        cv.imshow('butterfly', frame)
        if cv.waitKey(20) & 0xFF == ord('d'):
            break
    capture.release()
    cv.destroyAllWindows()


# Press the green button in the gutter to run the script.
if __name__ == '__main__':
    # show_grayscale('mei')
    show_butterfly()

# See PyCharm help at https://www.jetbrains.com/help/pycharm/
