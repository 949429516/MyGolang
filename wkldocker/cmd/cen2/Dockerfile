#docker 1.如果指定了其他命令，CMD可能会被忽略;2.若定义多个CMD，只会执行最后一个
FROM ubuntu
ENV name Docker
ENTRYPOINT echo "hello world"
#这个CMD不会执行
CMD echo "hello $name"
