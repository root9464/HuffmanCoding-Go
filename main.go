package main

import (
	"container/heap"
	"fmt"
)

type HuffmanTree interface {
	Freq() int
}

type HuffmanLeaf struct {
	freq  int
	value rune
}

type HuffmanNode struct {
	freq        int
	left, right HuffmanTree
}

func (self HuffmanLeaf) Freq() int {
	return self.freq
}

func (self HuffmanNode) Freq() int {
	return self.freq
}

type treeHeap []HuffmanTree

func (th treeHeap) Len() int { return len(th) }
func (th treeHeap) Less(i, j int) bool {
	return th[i].Freq() < th[j].Freq()
}
func (th *treeHeap) Push(ele interface{}) {
	*th = append(*th, ele.(HuffmanTree))
}
func (th *treeHeap) Pop() (popped interface{}) {
	popped = (*th)[len(*th)-1]
	*th = (*th)[:len(*th)-1]
	return
}
func (th treeHeap) Swap(i, j int) { th[i], th[j] = th[j], th[i] }

// Основная функция, которая строит дерево Хаффмана и печатает коды путем обхода
// построенное дерево Хаффмана

func buildTree(symFreqs map[rune]int) HuffmanTree {
	var trees treeHeap
	for c, f := range symFreqs {
		trees = append(trees, HuffmanLeaf{f, c})
	}
	heap.Init(&trees)
	for trees.Len() > 1 {
		// два дерева с наименьшей частотой

		a := heap.Pop(&trees).(HuffmanTree)
		b := heap.Pop(&trees).(HuffmanTree)

		// поместить в новый узел и повторно вставить в очередь
		heap.Push(&trees, HuffmanNode{a.Freq() + b.Freq(), a, b})
	}
	return heap.Pop(&trees).(HuffmanTree)
}

// Выводит коды Хаффмана из корня дерева Хаффмана. Он использует byte[] для
// хранения кодов
func printCodes(tree HuffmanTree, prefix []byte) {
	switch i := tree.(type) {
	case HuffmanLeaf:
		// Если это конечный узел, то он содержит один из входных
		// символы, выведите символ и его код из байта[]
		fmt.Printf("%c\t%d\t%s\n", i.value, i.freq, string(prefix))
	case HuffmanNode:
		// Присвоить 0 левому краю и повторить
		prefix = append(prefix, '0')
		printCodes(i.left, prefix)
		prefix = prefix[:len(prefix)-1]

		// Assign 1 to right edge and recur
		prefix = append(prefix, '1')
		printCodes(i.right, prefix)
		prefix = prefix[:len(prefix)-1]
	}
}

// Программа-драйвер для тестирования вышеуказанных функций
func main() {

	text := "hello world"

	symFreqs := make(map[rune]int)
	// прочитайте каждый символ и запишите частоты
	for _, c := range text {
		symFreqs[c]++
	}

	// пример дерева
	exampleTree := buildTree(symFreqs)

	// вывести результаты
	fmt.Println("Символ\tПовтор\tКодирование")
	printCodes(exampleTree, []byte{})
}
