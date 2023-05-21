package main
import "fmt"

func sort(arr []int) []int{
	if len(arr)<=1 {
		return arr
	}
	var middle int = len(arr)/2
	return merge(sort(arr[:middle]), sort(arr[middle:]))
}

func merge(arr1 []int, arr2 []int) (arr []int) {
	var idx1, idx2 int = 0, 0
	for (idx1 < len(arr1) && idx2 < len(arr2)) {
		if (arr1[idx1]>=arr2[idx2]) {
			arr = append(arr, arr2[idx2])
			idx2++
		} else {
			arr = append(arr, arr1[idx1])
			idx1++
		}
	}
	arr = append(arr, arr1[idx1:]...)
	arr = append(arr, arr2[idx2:]...)
	return
}

func main(){
	arr := []int{9,8,7,6,5,4,3,2,1}
	fmt.Println(sort(arr))
}
