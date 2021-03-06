# 贪心算法

## 概念

贪心的意思在于在作出选择时，每次都要选择对自身最为有利的结果，保证自身利益的最大化。贪心算法就是利用这种贪心思想而得出一种算法。 

贪心算法作为五大算法之一，在数据结构中的应用十分广泛。例如：在求最小生成树的 Prim 算法中，挑选的顶点是候选边中权值最小的边的一个端点。在 Kruskal 算法中，每次选取权值最小的边加入集合。在构造霍夫曼树的过程中也是每次选择最小权值的节点构造二叉树。这种每次在执行子问题的求解时，总是选择当前最优的情形，恰好符合贪心的含义。 

贪心算法可以简单描述为：大事化小，小事化了。对于一个较大的问题，通过找到与子问题的重叠，把复杂的问题划分为多个小问题。并且对于每个子问题的解进行选择，找出最优值，进行处理，再找出最优值，再处理。也就是说贪心算法是一种在每一步选择中都采取在当前状态下最好或最优的选择，从而希望得到结果是最好或最优的算法。

贪心算法在对问题求解时，总是做出在当前看来是最好的选择。也就是说，不从整体最优上加以考虑，所做出的仅是在某种意义上的**局部最优解**。贪心算法**不是对所有问题都能得到整体最优解**，但对范围相当广泛的许多问题他能产生**整体最优解**或者是**整体最优解的近似解**。

## 解题框架

**题目分析流程**

1. 建立数学模型来描述问题。
2. 把求解的问题分成若干个子问题。
3. 对每一子问题求解，得到子问题的局部最优解。
4. 把子问题的局部最优解合成原来问题的一个解。 

**伪代码**

```java
从问题的某一初始解出发
    while (能朝给定总目标前进一步) 
        do
            选择当前最优解作为可行解的一个解元素；
    由所有解元素组合成问题的一个可行解。
```

> 引用自：[五分钟了解一下什么是「贪心算法 」]( https://www.cxyxiaowu.com/852.html )

## 题目实战

### 1. 柠檬水找零

#### 题目描述

在柠檬水摊上，每一杯柠檬水的售价为 5 美元。

顾客排队购买你的产品，（按账单 bills 支付的顺序）一次购买一杯。

每位顾客只买一杯柠檬水，然后向你付 5 美元、10 美元或 20 美元。你必须给每个顾客正确找零，也就是说净交易是每位顾客向你支付 5 美元。

注意，一开始你手头没有任何零钱。

如果你能给每位顾客正确找零，返回 true ，否则返回 false 。

**示例 ：**

```
输入：[5,5,5,10,20]
输出：true
解释：
前 3 位顾客那里，我们按顺序收取 3 张 5 美元的钞票。
第 4 位顾客那里，我们收取一张 10 美元的钞票，并返还 5 美元。
第 5 位顾客那里，我们找还一张 10 美元的钞票和一张 5 美元的钞票。
由于所有客户都得到了正确的找零，所以我们输出 true。
```



#### 题目分析

**子问题：**

子问题很明显是给每位顾客找零

**局部最优解：**

对于顾客给的每一张钞票，找零时，优先从零钱中找到**可以用于找零的最大面额零钱**

例如，顾客给了一张20美元钞票，需要找零金额为15美元，我们优先考虑找`15 = 10 + 5`，如果没有10元钞票才考虑`15 = 5 + 5 + 5`

#### 代码实现

```go
func lemonadeChange(bills []int) bool {
    // 零钱, 值为该面额的纸币数, 20元无法用于找零
    change := [2]int{}
    // 用于标识change中对应的面值
    coins := [2]int{5, 10}
    for _, bill := range bills {
        coin := bill - 5 // 需要找零的钱
        // 选择当前最优解
        // 每次选可以给的最大零钱
        indx := 1
        for coin > 0 && indx >= 0 {
            if change[indx] > 0 && coin >= coins[indx] {
                coin -= coins[indx]
                change[indx]--
            } else {
                indx--
            }
        }
        // 没有找到合适的零钱
        if indx < 0 {
            return false
        }
        // 将支付的钱放进零钱袋
        switch bill {
        case 10:
            change[1]++
        case 5:
            change[0]++
        }
    }

    return true
}
```

### 2. 跳跃游戏

#### 题目描述

给定一个非负整数数组，你最初位于数组的第一个位置。

数组中的每个元素代表你在该位置可以跳跃的最大长度。

判断你是否能够到达最后一个位置。

**示例 :**

```
输入: [2,3,1,1,4]
输出: true
解释: 我们可以先跳 1 步，从位置 0 到达 位置 1, 然后再从位置 1 跳 3 步到达最后一个位置。
```



#### 题目分析

**子问题：**

对于任意位置`i`，它所能跳到的范围为`[i, i + nums[i]]`，设目标位置为y，如果`i + nums[i] >= y`，则说明位置`i`能够跳到位置`y`，如果`y`为数组长度，即最后一个位置，那么`i`能够跳到最后。

**局部最优解：**

从后往前看，如果位置`j`能够到达最后一个位置，位置`i`能够到达位置`j`，则说明位置`i`能够到达最后一个位置。因此只需要维护一个能够跳到最后一个位置的最小下标`minIndx`即可，每当找到一个位置`i`，使得能够跳到位置`minIndx`，则更新`minIndx`

**合并局部最优解**

如果`minIndx == 0`，则认为从数组起始位置能够跳到最后一个位置

#### 代码实现

```go
func canJump(nums []int) bool {
	// 从右往左遍历，判断当前位置能不能跳到最后
	minIndx := len(nums)-1 // 能够跳到最后的最小下标
	for i := len(nums)-2; i >= 0; i-- {
		if nums[i] + i >= minIndx {
			minIndx = i
			continue
		}
	}
	return minIndx == 0
}
```

