public class Insertion_Sort {
    public int[] sortArray(int[] nums) {
        // Loop through the array starting from the second element
        for (int i = 1; i < nums.length; i++) {
            int key = nums[i]; // The element to be inserted
            int j = i - 1;

            // Move elements that are greater than key to one position ahead
            while (j >= 0 && nums[j] > key) {
                nums[j + 1] = nums[j];
                j--;
            }
            // Insert the key at the correct position
            nums[j + 1] = key;
        }
        return nums; // Return the sorted array
    }
}
