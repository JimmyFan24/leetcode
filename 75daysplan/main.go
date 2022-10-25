package main

import (
	"fmt"
	"leetcode/75daysplan/level8"
)

func main() {

	t := level8.SearchMatrix([][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}, 5)
	fmt.Println(t)
	//level8.MaximalSquare([][]string{{"0"}}}
	//level8.MaximalSquare([][]byte{{'1','0'},{'0','1'}})
	//level8.GetHappyString(3,9)
	//level8.CreateBinaryTree([][]int{{20,15,1},{20,17,0},{50,20,1},{50,80,0},{80,19,1}})
	//ans := 3
	//fmt.Println(ans >> 1)
	//level8.ConstructFromPrePost([]int{1,2,4,5,3,6,7},[]int{4,5,2,6,7,3,1})
	//level8.BuildArray([]int{1,2},4)
	//level8.SuggestedProducts([]string{"mobile", "mouse", "moneypot", "monitor", "mousepad"}, "mouse")
	//level8.Evaluate("(name)is(age)yearsold",[][]string{{"name","bob"},{"age","two"}})
	//level8.DailyTemperatures([]int{89,62,70,58,47,47,46,76,100,70})
	/*c := level7.Constructor1472("aaa")
	c.Visit("bbb")
	c.Visit("ccc")
	c.Visit("ddd")
	f := c.Back(1)
	g := c.Back(1)
	t := c.Forward(1)
	c.Visit("eee")
	r := c.Forward(2)
	m := c.Back(2)
	u := c.Back(7)
	fmt.Println(f, g, t, r, m, u)*/
	//level7.NumSpecialEquivGroups([]string{"abcd","cdab","cbad","xyzz","zzxy","zzyx"})
	//level7.CustomSortString("cba", "abcd")
	//level7.GetSmallestString1(96014,2095650)
	/*c := level7.Constructor("abc", 2)
	c.Next()
	c.HasNext()
	c.Next()
	c.HasNext()*/
	//level7.SpiralMatrixIII(5,6,1,4)
	//level7.DisplayTable([][]string{{"David","3","Ceviche"},{"Corina","10","Beef Burrito"},{"David","3","Fried Chicken"},{"Carla","5","Water"},{"Carla","5","Ceviche"},{"Rous","3","Ceviche"}})
	//level7.SmallestNumber("IIIDIDDD")
	//level7.PartitionString("abacaba")
	/*node13 := &level7.ListNode{Val: 0,}
	node12 := &level7.ListNode{Val: 5,Next: node13}
	node11 := &level7.ListNode{Val: 5,Next: node12}
	node10 := &level7.ListNode{Val: 2,Next: node11}
	node9:= &level7.ListNode{Val: 4,Next: node10}
	node8 := &level7.ListNode{Val: 9,Next: node9}
	node7 := &level7.ListNode{Val: 7,Next: node8}
	node6 := &level7.ListNode{Val: 1,Next: node7}
	node5 := &level7.ListNode{Val: 8,Next: node6}
	node4 := &level7.ListNode{Val: 6,Next: node5}
	node3 := &level7.ListNode{Val: 2,Next: node4}
	node2 := &level7.ListNode{Val: 0,Next: node3}
	node1 := &level7.ListNode{Val: 3,Next: node2}
	level7.SpiralMatrix(3,5,node1)*/
	//level7.PathInZigZagTree(14)
	//level7.MatrixScore([][]int{{0,1},{1,1}})
	//level7.CountMaxOrSubsets([]int{2, 2, 2})
	//level7.CountTriplets([]int{2,3,1,6,7})
	//node2 :=&level7.ListNode{Val: 90}
	//node1 := &level7.ListNode{Val: 100,Next: node2}
	//level7.SwapNodes(node1,2)
	//level7.NumOfSubarrays([]int{11,13,17,23,29,31,7,5,2,3},3,5)
	//level7.CountSubIslands([][]int{{1,1,1,0,0},{0,1,1,1,1},{0,0,0,0,0},{1,0,0,0,0},{1,1,0,1,1}},[][]int{{1,1,1,0,0},{0,0,1,1,1},{0,1,0,0,0},{1,0,1,1,0},{0,1,0,1,0}})
	//level7.MinSwaps("]]][[[")
	//level7.FrequencySort("tree")
	//level7.KWeakestRows([][]int{{1, 1, 0, 0, 0}, {1, 1, 1, 1, 0}, {1, 0, 0, 0, 0}, {1, 1, 0, 0, 0}, {1, 1, 1, 1, 1}}, 3)
	//level7.AllCellsDistOrder(1,2,0,0)
	//level7.GreatestLetter("arRAzFif")
	//level7.NumRookCaptures([][]byte{{'.','.','.','.','.','.','.','.'},{'.','.','.','p','.','.','.',',','.'},{'.','.','.','R','.','.','.','p'},{'.','.','.','.','.','.','.','.'},{'.','.','.','.','.','.','.','.'},{'.','.','.','p','.','.','.','.'},{'.','.','.','.','.','.','.','.'},{'.','.','.','.','.','.','.','.'}})
	//level7.NumTilePossibilities("AAB")
	//level7.MaximumXOR([]int{3,2,4,6})
	//level7.CountHillValley([]int{6,6,5,5,4,1})
	//level7.DivisorSubstrings(3003,3)
	//level7.FindPoisonedDuration([]int{1,2 },2)
	//level7.MakeGood("abBAcC")
	//level7.IsThree(14)
	//level7.CountTriples(5)
	//level7.CountStudents([]int{1,1,1,0,0,1},[]int{1,0,0,0,1,1})
	//level7.CheckXMatrix([][]int{{2,0,0,1},{0,3,1,0},{0,5,2,0},{4,0,0,2}})
	//level7.MinStartValue([]int{1,-2,-3})
	//level7.FindWords([]string{"Hello","Alaska","Dad","Peace"})
	//level7.IslandPerimeter([][]int{{0,1,0,0},{1,1,1,0},{0,1,0,0},{1,1,0,0}})
	//level7.Intersection([][]int{{3,1,2,4,5},{1,2,3,4},{3,4,5,6}})
	//level7.MinimumAbsDifference([]int{4,2,1,3})
	//level7.CheckDistances("abaccb",[]int{1,3,0,5,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0})
	//level7.ShortestToChar("loveleetcode",'e')
	//level7.MinTimeToType("zjpc")
	//level7.SortByBits([]int{0,1,2,3,4,5,6,7,8})
	//level7.OddCells(2, 3, [][]int{{0, 1}, {1, 1}})
	//level7.MinTimeToVisitAllPoints([][]int{{1,1},{3,4},{-1,0}})
	//level7.FreqAlphabets("10#11#12")
	//level7.ReplaceDigits("a1c1e1")
	//level7.RemoveOuterParentheses("(()())(())")
	//level7.FlipAndInvertImage([][]int{{1,1,0},{1,0,1},{0,0,0}})
	//level6.NumSubmatrixSumTarget([][]int{{0,1,0},{1,1,1},{0,1,0}},0)
	//level6.MinNumberOperations([]int{1,2,3,2,1})
	//level6.CanSeePersonsCount([]int{10, 6, 8, 5, 11, 9})
	//level6.MaxSatisfaction([]int{4,3,2})
	//level6.ConcatenatedBinary(42)
	//level6.MaxScoreWords([]string{"ac", "abd", "db", "ba", "dddd", "bca"}, []byte{'a', 'a', 'a', 'b', 'b', 'b', 'c', 'c', 'd', 'd', 'd', 'd'}, []int{6, 4, 4, 7, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
	//level6.UniquePathsIII([][]int{{0,1},{2,0}})
	//level6.LargestDivisibleSubset([]int{5,9,18,54,108,540,90,180,360,720})

	//level6.SumEvenAfterQueries([]int{1, 2, 3, 4}, [][]int{{1, 0}, {-3, 1}, {-4, 0}, {2, 3}})

	//level6.MaxProfit([]int{6, 1, 3, 2, 4, 7})
	//level6.CountPrimes(10)
	//level6.PermuteUnique([]int{3,4,5})
	//level6.FindDuplicate([]string{"root/a 1.txt(abcd) 2.txt(efgh)","root/c 3.txt(abcd)","root/c/d 4.txt(efgh)","root 4.txt(efgh)"})
	//level6.MaxNumberOfFamilies(4, [][]int{{4, 3}, {1, 4}, {4, 6}, {1, 7}})
	//level6.NumDecodings("12")
	//level6.SubsetsWithDup([]int{1,2,2})
	//level6.SearchMatrix([][]int{{1,3,5,7},{10,11,16,20},{23,30,34,60}},3)
	//level6.TrailingZeroes(30)
	//fmt.Println(10*9*8*7*6*5*4*3*2*1)
	//level6.EvalRPN([]string{"3","-4","+"})
	/*t := level6.Constructor355()
	t.PostTweet(2,5)
	t.Follow(1,2)
	t.Follow(1,2)
	text1 := t.GetNewsFeed(1)
	t.Follow(2,1)
	text2 := t.GetNewsFeed(2)
	t.Unfollow(2,1)
	text3 := t.GetNewsFeed(2)
	fmt.Println(text1,text2,text3)*/
	//level6.IntegerBreak(10)
	//level6.CountNumbersWithUniqueDigits(3)
	//level6.HIndex([]int{1,3,1})
	//level6.NthUglyNumber(10)
	//level5.WordBreak("", []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"})
	//level5.ReverseWords("  hello world  ")
	//level5.MinimumTotal([][]int{{2},{3,4},{6,5,7},{4,1,8,3}})
	//level5.RestoreIpAddresses("25525511135")
	/*l6 := &level5.ListNode{
		6, nil,
	}

	l5 := &level5.ListNode{
		5,
		l6,
	}
	l4 := &level5.ListNode{
		4,
		l5,
	}
	l3 := &level5.ListNode{
		3,
		l4,
	}
	l2 := &level5.ListNode{
		2,
		l3,
	}
	l1 := &level5.ListNode{
		1,
		l2,
	}
	level5.ReverseBetween(l1,1,2)*/
	//level5.Partition(l1, 3)
	//level5.Combine(4,2)
	//level5.Insert([][]int{{1,5}},[]int{2,3})
	//
	//level5.Insert([][]int{{1,2},{3,5},{6,7},{8,10},{12,16}},[]int{4,8})
	//level5.GenerateMatrix(6)
	//level5.CombinationSum2([]int{10,1,2,7,6,1,5},8)
	//level5.Multiply("498828660196","840477629533")
	//fmt.Println(498828660196*840477629533)
	//level5.NextPermutation([]int{3,2,1})
	//level5.IntToRoman(1994)
	//level5.FindFarmland([][]int{{1, 1}, {0, 0}})
	//level5.MinFallingPathSum([][]int{{-84,-36,2},{87,-79,10},{42,10,63}})
	//level5.HasAlternatingBits(5)
	//level5.PartitionArray([]int{3, 6, 1, 2, 5}, 2)
	//level5.CanVisitAllRooms([][]int{{2,3},{},{2},{1,3}})
	//level5.MctFromLeafValues([]int{6, 2, 4})

	/*node7 := &level5.TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	}
	node8 := &level5.TreeNode{
		Val:   4,
		Left:  nil,
		Right: nil,
	}
	node6 := &level5.TreeNode{
		Val:   7,
		Left:  nil,
		Right: nil,
	}
	node5 := &level5.TreeNode{
		Val:   13,
		Left:  nil,
		Right: nil,
	}*/
	/*node4 := &level5.TreeNode{
		Val:  2,
		Left:  nil,
		Right: nil,
	}*/
	/*node3 := &level5.TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	}
	node2 := &level5.TreeNode{
		Val:   3,
		Left:  nil,
		Right: node3,
	}
	node1 := &level5.TreeNode{
		Val:   1,
		Left:  node2,
		Right: nil,
	}
	level5.RecoverTree(node1)*/
	//level5.PathSum(node1,22)*/
	//level5.SumNumbers(node1)
	//level5.SubtreeWithAllDeepest(node1)
	//level5.RemoveLeafNodes(node1, 3)
	//level5.LetterCasePermutation("a1b2")
	//level5.RemoveOccurrences("yjyjqnaxlbqnaxlbfss","yjqnaxlb")
	//level5.CountSquares([][]int{{1,0,1},{1,1,0},{1,1,0}})
	//level5.GarbageCollection([]string{"MMM","PGM","GP"},[]int{3,10})
	//level5.HasAlternatingBits(5)

	//level5.MinimumOperations([]int{1,5,0,3,5})
	//level5.CanBeEqual([]int{3,9,7},[]int{3,9,11})
	//level5.SpecialArray([]int{3,9,7,8,3,8,6,6})
	//level5.MakeEqual([]string{"abc","aabc","bc"})
	//level4.CommonChars([]string{"ccbdbcba", "cbdda", "accdcdbb"})
	//level4.NumUniqueEmails([]string{"a@leetcode.com", "b@leetcode.com", "c@leetcode.com"})
	//level4.RelativeSortArray([]int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19}, []int{2, 1, 4, 3, 9, 6})
	//level4.FindMiddleIndex([]int{0, 0, 0, 0})
	//level4.MinCostToMoveChips([]int{2, 2, 2, 3, 3})
	//level4.SortString("rat")
	//level4.FrequencySort([]int{-1, 1, -6, 4, 5, -6, 1, 4, 1})
	//level4.CountBinarySubstrings("1100")
	//level4.FindShortestSubArray([]int{2, 1})
	/*kStream := level4.Constructor703(3, []int{4, 5, 8, 2})
	a := kStream.Add(3)
	b := kStream.Add(5)
	c := kStream.Add(10)
	d := kStream.Add(9)
	e := kStream.Add(4)
	fmt.Println(a, b, c, d, e)*/
	//level4.ShortestCompletingWord("1s3 PSt", []string{"step", "steps", "stripe", "stepple"})
	//level4.MinAddToMakeValid("())")
	//fmt.Println(((2<<2)-1))
	//fmt.Println(2^3^4^7^5)
	//level4.GetMaximumXor([]int{0},1)
	/*node7 := &level4.TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}
	node6 := &level4.TreeNode{
		Val:   0,
		Left:  nil,
		Right: nil,
	}
	node5 := &level4.TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}
	node4 := &level4.TreeNode{
		Val:   0,
		Left:  nil,
		Right: nil,
	}
	node3 := &level4.TreeNode{
		Val:   1,
		Left:  node6,
		Right: node7,
	}
	node2 := &level4.TreeNode{
		Val:   0,
		Left:  node4,
		Right: node5,
	}*/
	/*node1 := &level4.TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}
	level4.SumRootToLeaf(node1)*/
	/*level4.DelNodes(node1,[]int{3,5})*/
	//level4.FindAndReplacePattern([]string{"deq", "mee", "aqq", "dkd", "ccc"}, "abb")
	//level4.CheckArithmeticSubarrays([]int{4,6,5,9,3,7},[]int{0,0,2},[]int{2,3,5})
	//level4.DiagonalSort([][]int{{3,3,1,1},{2,2,1,2},{1,1,1,2}})
	//level4.ExecuteInstructions(3,[]int{0,1},"RRDDLU")
	//level4.MinOperations("001011")
	//level4.MaxIncreaseKeepingSkyline([][]int{{3,0,8,4},{2,4,5,7},{9,2,6,3},{0,3,1,0}})
	/*sub := level4.Constructor1476([][]int{{1,2,1},{4,3,4},{3,2,1},{1,1,1}})
	sub.GetValue(0,2)*/
	//level4.CreateBinaryTree([][]int{{20, 15, 1}, {20, 17, 0}, {50, 20, 1}, {50, 80, 0}, {80, 19, 1}})
	//level4.WateringPlants([]int{2,2,3,3},5)
	//level4.MinPairSum([]int{3,5,4,2,4,6})
	//level4.ProcessQueries([]int{7,5,5,8,3},8)
	//level4.GroupThePeople([]int{3,3,3,3,3,1,3},)
	//level4.MaxSumAfterPartitioning([]int{1,15,7,9,2,5,10},3)
	//level4.NumberOfBeams([]string{"011001", "000000","010100", "001000"})
	//level4.IntervalIntersection([][]int{{3, 5}, {9, 20}}, [][]int{{4, 5}, {7, 10}, {11, 12}, {14, 15}, {16, 20}})
	//level4.IntervalIntersection([][]int{{0,2},{5,10},{13,23},{24,25}},[][]int{{1,5},{8,12},{15,24},{25,26}})
	//level4.ComplexNumberMultiply("78+-76i","-86+72i")
	//level4.MinSteps("cotxazilut","nahrrmcchxwrieqqdwdpneitkxgnt")
	/*sm := level4.NewConstructor()
	sm.AddBack(32)
	a := sm.PopSmallest()
	b:= sm.PopSmallest()
	c := sm.PopSmallest()
	sm.AddBack(1)
	fmt.Println(a,b,c)*/
	//level4.QueensAttacktheKing([][]int{{0,0},{1,1},{2,2},{3,4},{3,5},{4,4},{4,5}},[]int{3,3})
	//level4.MemLeak(2,2)
	/*code := level4.URLConstructor()
	shorturl := code.Encode("https://baidu.com/1udbcdc")
	url := code.Decode(shorturl)
	fmt.Println(url)*/
	/*node8 := &level4.ListNode{
		Val:  0,
		Next: nil,
	}
	node7 := &level4.ListNode{
		Val:  2,
		Next: node8,
	}
	node6 := &level4.ListNode{
		Val:  5,
		Next: node7,
	}
	node5 := &level4.ListNode{
		Val:  0,
		Next: node6,
	}
	node4 := &level4.ListNode{
		Val:  1,
		Next: node5,
	}
	node3 := &level4.ListNode{
		Val:  3,
		Next: node4,
	}
	node2 := &level4.ListNode{
		Val:  1,
		Next: node3,
	}/*
	level4.DelNodes(node2,[]int{3,5})
	//level4.LongestNiceSubstring("cChH")
	//level4.CountGoodTriplets([]int{3,0,1,1,9,7},7,2,3)
	//level4.MinBitFlips(10,82)
	//level4.MaxDepth("(1+(2*3)+((8)/4))+1")
	/*os := level4.CConstructor(5)
	os.Insert(3,"ccccccc")
	os.Insert(1,"aaaaaaa")
	os.Insert(2,"bbbbbbb")
	os.Insert(5,"eeeeeee")
	os.Insert(4,"ddddddd")*/
	//level4.MinSubsequence([]int{10, 2, 5})
	//level4.CountBalls(11, 104)
	/*c := level4.Constructor()
	c1 := c.Ping(1)
	c2 := c.Ping(100)
	c3 := c.Ping(3001)
	c4 := c.Ping(3002)
	fmt.Println(c1, c2, c3, c4)*/
	//level4.SmallestEqual([]int{0,1,2})
	//level4.TwoOutOfThree([]int{1,1,3,2}, []int{2,3}, []int{3})
	//level4.CalPoints([]string{"5","2","C","D","+"})
	//level4.IsSumEqual("abc", "cba", "cdb")
	//level4.Maximum69Number(9969)
	//level4.DiagonalSum([][]int{{1,2,3},{4,5,6},{7,8,9}})
	//level4.NumOfStrings([]string{"a","abc","bc","d"},"abc")
	//level3.DivideString("abcdefghij", 3, 'x')
	//level3.ConvertTime("09:41","10:34")
	//level3.TrimMean([]int{1,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,3})
	//level3.MaxDistance([]int{1,1,1,6,1,1,1})
	//level3.DuplicateZeros([]int{1,0,2,3,0,4,5,9})
	//level3.GreatestLetter( "AAAA")
	//level3.FindDifference([]int{1,2,3,3},[]int{1,1,2,2})
	//level3.ReplaceElements([]int{17,18,5,4,6,1})
	//level3.MinOperations([]int{1,5,2,4,1})
	//level3.FinalPrices([]int{8,4,6,2,3})
	//level3.CountGoodSubstrings("xyzzaz")
	//level3.MaximumUnits([][]int{{2,1},{4,4},{3,1},{4,1},{2,4},{3,4},{1,3},{4,3},{5,3},{5,3}},13)
	//level3.LargestAltitude([]int{-4,-3,-2,-1,4,3,2})
	//level3.CountLargestGroup(13)
	//level3.RepeatedCharacter("abccbaacz")
	//level3.CountGoodTriplets([]int{3,0,1,1,9,7},7,2,3)
	//level3.LastStoneWeight([]int{2,7,4,1,8,1})
	//level3.RemoveDuplicates("aababaab")
	//level3.SubsetXORSum([]int{5, 1, 6})
	//level3.RemovePalindromeSub("ababb")
	//level3.MergeAlternately("ab", "nba")
	//level3.CountOperations(2, 3)
	//level3.NumberOfPairs([]int{1, 3, 2, 1, 3, 2, 2})
	//level3.SumZero(4)
	//level3.CountPoints("B0R0G0R9R0B0G0")
	//level3.LargestLocal([][]int{{9, 9, 8, 1}, {5, 6, 2, 6}, {8, 2, 6, 4}, {6, 2, 2, 2}})
	//level3.UniqueMorseRepresentations([]string{"gin", "zen", "gig", "msg"})
	//level3.CountConsistentStrings("ab", []string{"ad", "bd", "aaab", "baa", "badab"})
	//level3.CellsInRange("K1:L2")
	//level3.SumOddLengthSubarrays([]int{1, 4, 2, 5, 3})
	//level3.DecodeMessage("eljuxhpwnyrdgtqkviszcfmabo", "vkbs bs t suepuv")
	//level3.ArithmeticTriplets([]int{4, 5, 6, 7, 8, 9}, 2)
	//tmp := level3.FindLUSlength("ab", "cc")
	//level3.NextGreaterElement([]int{4, 1, 2}, []int{1, 4, 3, 2})
	//tmp := level3.ConvertToBase7(-7)
	//fmt.Println(tmp)
	//level3.FindMaxConsecutiveOnes([]int{1,1,0,1,1,1})
	//"2-5g-3-J" "2-5G-3J" k==2
	//"2-4A0r7-4k" "24A0-R74K" k==4
	//level3.LicenseKeyFormatting("VDxVMWm-K-U-xosBSoNOGdKjFu-sgStQBE-NNqFLLUWuIIuyTInpNIuGoxxKvlTzkuadwlb-UUs-xgwxKxRIb-qX-FdWMYeVU-CzxIrpFapDTQ",22)
	//level3.FindComplement(1)
	//level3.Gcd(4,6)
	//level3.FourSum([]int{-3,-1,0,2,4,5},0)
	//level3.ThreeSumClosest([]int{-100,-98,-2,-1},-101)
	/*res := int64(7)
	fmt.Println(strconv.FormatInt(res, 2))
	fmt.Println(strconv.FormatInt(-res, 2))
	fmt.Println(strconv.FormatInt(res&(-res), 10))*/
	//level3.SingleNumberIII([]int{1,2,1,2,3,4})
	//level3.ReverseStr("abcdefg",3)
	//level3.ReverseWords("Let's take LeetCode contest")
	//level3.CheckPerfectNumber(2016)
	//level3.MaxProduct([]int{-2,0,-1})
	//level3.Rob([]int{1,2,3,1})
	//level3.Jump([]int{4,1,1,3,1,1,1})
	//level3.PartitionLabels("ababcbacadefegdehijhklij")
	//level3.LongestCommonSubsequence("oxcpqrsvwf","shmtulqrypy")
	//level3.IsMonotonic([]int{1,2,3,3,4,5})
	//level3.Transpose([][]int{{1,2,3},{4,5,6},{7,8,9}})
	//level3.ToGoatLatin("I speak Goat Latin")
	//level3.SortArrayByParityII([]int{4,5,2,7})
	//level3.BuddyStrings("abac","abad")
	//[3,5,1,6,2,9,8,null,null,7,4]
	/*
		node9 := &level2.TreeNode{
			Val:   9,
			Left:  nil,
			Right: nil,
		}
		node8 := &level2.TreeNode{
			Val:   7,
			Left:  nil,
			Right: nil,
		}*/
	/*node7 := &level2.TreeNode{
		Val:   8,
		Left:  node8,
		Right: node9,
	}*/

	/*node6 := &level3.TreeNode{
		Val:   7,
		Left:  nil,
		Right: nil,
	}
	node5 := &level3.TreeNode{
		Val:   5,
		Left:  nil,
		Right: nil,
	}
	node4 := &level3.TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}
	node3 := &level3.TreeNode{
		Val:   9,
		Left:  nil,
		Right: node6,
	}
	node2 := &level3.TreeNode{
		Val:   2,
		Left:  node4,
		Right: node5,
	}
	node1 := &level3.TreeNode{
		Val:   4,
		Left:  node2,
		Right: node3,
	}
	level3.FindTilt(node1)*/
	//level3.IncreasingBST(node1)
	//level3.LeafSimilar(node1,nil)
	//level3.CountSegments("                ")
	//level3.CanJump([]int{2,1,0,0})
	//level3.MinPathSum([][]int{{1,3,1},{1,5,1},{4,2,1}})
	//level3.LargestNumber([]int{3,30,34,5,9})
	//fmt.Println("344">"390")
	//level3.IsPowerOfFour(17)
	//level3.WordPattern("abba","dog dog dog dog")
	//level3.AddDigits(38)
	//level3.SummaryRanges([]int{0,1,2,4,5,7})
	//res := 'A' +1
	//fmt.Println(1%26)
	//level3.ConvertToTitle(701)
	//level3.ContainsNearbyDuplicate([]int{1,0,1,1},1)
	//r,_ := strconv.Atoi(string(res)) fmt.Println(string(res))
	//level2.FindMedianSortedArrays([]int{1,2},[]int{3,4})
	//level2.Rotate([][]int{{1,2,3},{4,5,6},{7,8,9}})
	//level2.IsMatch("mississippi","mis*is*ip*.")
	//level2.CombinationSum4([]int{1,2,3},4)
	//level2.NewKthSmallest([][]int{[]int{1,5,9},[]int{10,11,13},[]int{12,13,15}},8)
	//level2.MyNewPow(2.0000,0)
	//level2.Divide(-2147483648,1)
	//fmt.Println( math.MaxInt32)
	//level2.Search([]int{1},1)
	//level2.SearchRange([]int{1,4},4)
	//nums := []int{0,0,4,1,2,1,2}
	//newNums := level1.RunningSum(nums)
	//index := level1.PivotIndex(nums)
	//s := "acb"
	//t := "ahbgdc"
	//level1.IsSubsequence(s,t)
	/*node5 := &level2.ListNode{
		Val:  5,
		Next: nil,
	}
	node4 := &level2.ListNode{
		Val:  4,
		Next: node5,
	}
	node3 := &level2.ListNode{
		Val:  3,
		Next: node4,
	}
	node2 := &level2.ListNode{
		Val:  2,
		Next: node3,
	}
	node1 := &level2.ListNode{
		Val:  1,
		Next: node2,
	}*/
	//level2.Flatten(node1)
	//level1.ReverseList(node1)
	//level1.MergeTwoLists(node1,node3)

	//level1.IsInterleave("","b","b")
	//level1.LongestConsecutive(nums)
	//level1.KidsWithCandies(nums,3)
	//level1.Makesquare(nums)
	//level1.SingleNumber(nums)
	//level1.Generate(5)
	//level1.HammingWeight(001011)
	//level1.MoveZeroes(nums)
	//	res := []byte("ZA")
	//fmt.Println(res)
	//titleToNumber("BBA")
	//backspaceCompare("a##c","#c#d")
	/*image := [][]int{[]int{1,1,1},[]int{1,1,0},[]int{1,0,1}}
	sr := 1
	sc := 1
	color := 2
	level1.FloodFill(image,sr,sc,color)*/
	//level1.DecodeString("3[2[a]]")
	//fmt.Println(11/10);
	/*board := [][]byte{
		{'A','B','C','E'},
		{'S','F','C','S'},
		{'A','D','E','E'},
	}*/
	//level1.LongestCommonPrefix([]string{"reflower","flow","flight"})
	//level1.IsValid("()")
	//nums := []int{7,2,8,5,0,9,1,2,9,5,3,6,6,7,3,2,8,4,3,7,9,5,7,7,4,7,4,9,4,7,0,1,1,1,7,4,0,0,6}
	//level1.PlusOne(nums)
	//level1.IsPalindrome(node1)
	//level1.Exist(board,"ABCC")
	/*s := "   fly me to the moon  "
	reg := regexp.MustCompile(`[A-Za-z]+`)
	str := reg.FindAllString(s,-1)

	fmt.Println(str)*/
	//level1.ArrangeCoins(180)
	//level1.ThirdMax([]int{1,2,3})
	//level1.LongestPalindrome("aaaa")
	/*s := "qlhxagxdqh"

	sl := []string{"qlhxagxdq","qlhxagxdq","lhyiftwtut","yfzwraahab"}
	level2.NumMatchingSubseq(s,sl)*/
	//level2.MyAtoi("-5-")
	//level2.ThreeSum([]int{-1,0,1,2,-1,-4})
	//fmt.Println((1<<31))
	//level2.MaxProfit([]int{7,1,5,3,4,6})
	//res := level2.LetterCombinations("2")
	//fmt.Println(res)
	//level2.GenerateParenthesis(2)
	//level2.RemoveNthFromEnd(node1,2)
	//res := level2.Subsets([]int{1,2,4})
	//fmt.Println(res)
	//level2.FindKthLargest([]int{3,2,1,5,6,4},2)
	//level2.TopKFrequent([]int{1,3},2)
	//level2.OddEvenList(node1)
	//level2.GetSum(1,2)
	//level2.LengthOfLIS([]int{10,9,2,5,3,7,101,18})
	/*t3 := &level2.TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}
	t2 := &level2.TreeNode{
		Val:   2,
		Left:  t3,
		Right: nil,
	}

	t1 := &level2.TreeNode{
		Val:   1,
		Left:  nil,
		Right: t2,
	}*/

	//level2.Flatten(t1)
	//level2.ThreeSum([]int{-1,0,1,2,-1,-4})
	/*node7 := &level2.Node{
		Val:   7,
		Left:  nil,
		Right: nil,
		Next:  nil,
	}
	node6 := &level2.Node{
		Val:   6,
		Left:  nil,
		Right: nil,
		Next:  nil,
	}
	node5 := &level2.Node{
		Val:   5,
		Left:  nil,
		Right: nil,
		Next:  nil,
	}
	node4 := &level2.Node{
		Val:   4,
		Left:  nil,
		Right: nil,
		Next:  nil,
	}
	node3 := &level2.Node{
		Val:   3,
		Left:  node6,
		Right: node7,
		Next:  nil,
	}
	node2 := &level2.Node{
		Val:   2,
		Left:  node4,
		Right: node5,
		Next:  nil,
	}
	node1 := &level2.Node{
		Val:   1,
		Left:  node2,
		Right: node3,
		Next:  nil,
	}*/

	//str := level2.CountAndSay(6)
	//fmt.Println(str)
	//nu := [][]int{{0,1,0},{0,0,1},{1,1,1},{0,0,0}}
	//level2.Connect(node1)
	//level2.GameOfLife(nu)
	//level2.AddStrings("9","9")
	//res := level2.CombinationSum([]int{1,1},2)
	//fmt.Println(res)
	//level2.AddBinary("1010","1011")
	//level2.MinimumSum(2932)
	//level2.SubtractProductAndSum(234)
	//level2.CreateTargetArray([]int{0,1,2,3,4},[]int{0,1,2,2,1})

}
func titleToNumber(columnTitle string) int {

	res, val := 0, 0

	for _, v := range columnTitle {
		val = int(v - 'A' + 1)
		res = res*26 + val
	}
	return res
}
func backspaceCompare(s string, t string) bool {
	sAfter := ""
	tAfter := ""

	for _, b := range s {
		if string(b) == "#" {
			sAfter = sAfter[:len(sAfter)-1]
		} else {
			sAfter += string(b)
		}

	}

	for _, b := range t {
		if string(b) == "#" {
			tAfter = tAfter[:len(tAfter)-1]
		} else {
			tAfter += string(b)
		}

	}
	if sAfter == tAfter {
		return true
	}
	return false
}
