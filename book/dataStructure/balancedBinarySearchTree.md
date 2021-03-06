# 平衡二叉搜索树

## 定义

**平衡二叉树定义：**

二叉树中任意一个结点的左右子树的高度相差不能大于1。从这个定义上来看，完全二叉树、满二叉树其实都是平衡二叉树，但非完全二叉树也有可能是平衡二叉树。

**平衡二叉搜索树：**不仅满足上面平衡二叉树的定义，还满足二叉搜索树的特点。最先被发明的平衡二叉搜索树是[AVL树](https://zh.wikipedia.org/wiki/AVL%E6%A0%91)，它严格符合平衡二叉搜索树的定义，是一种高度平衡的二叉搜索树。



## 设计的取舍

很多平衡二叉搜索树其实并没有严格符合上面的定义（树中任意一个结点的左右子树的高度相差不能大于1），比如红黑树，它从根节点到各个叶子节点的最长路径，有可能会比最短路径大一倍。

学习数据结构与算法是为了应用到实际开发中，所以，没有必要死抠定义。要从这个数据结构的由来，去理解**“平衡”**的意思。

发明平衡二叉查找树这类数据结构的**初衷是**，**解决**普通二叉查找树在频繁的插入、删除等动态更新的情况下，出现**时间复杂度退化**的问题。 

所以，**平衡二叉查找树中“平衡”的意思，其实就是让整棵树左右看起来比较“对称”、比较“平衡”，不要出现左子树很高、右子树很矮的情况。这样就能让整棵树的高度相对来说低一些，相应的插入、删除、查找等操作的效率高一些。** 

所以，如果我们现在设计一个新的平衡二叉查找树，**只要树的高度不比 log2n 大很多**（比如树的高度仍然是对数量级的），尽管它不符合我们前面讲的严格的平衡二叉查找树的定义，但我们仍然可以说，这是一个合格的平衡二叉查找树。 

**工程中为什么大家都喜欢用红黑树这种平衡二叉搜索树？**

Treap、Splay Tree，绝大部分情况下，它们的操作效率都很高，但是也无法避免极端情况下时间复杂度的退化。尽管这种几率不大，但是对于单次操作时间非常敏感的场景来说，它们并不适用。

AVL 树是一种高度平衡的二叉树，所以查找的效率非常高，但是，有利就有弊，AVL 树为了维持这种高度的平衡，就要付出更多的代价。每次插入、删除都要做调整，就比较复杂、耗时。所以，对于有频繁的插入、删除操作的数据集合，使用 AVL 树的代价就有点高了。

红黑树只是做到了近似平衡，并不是严格的平衡，所以在维护平衡的成本上，要比 AVL 树要低。

所以，红黑树的插入、删除、查找各种操作性能都比较稳定。对于工程应用来说，要面对各种异常情况，为了支撑这种工业级的应用，我们更倾向于这种性能稳定的平衡二叉查找树。

## 平衡二叉搜索树实现

- AVL
- [红黑树（Red Black Tree）](./redBlackTree.md)
- 伸展树（Splay Tree）
- 树堆（Treap）



## 引用

> [数据结构与算法之美]( https://time.geekbang.org/column/intro/100017301 )