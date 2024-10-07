public class Solution {
    public IList<IList<int>> Subsets(int[] nums) {
        IList<IList<int>> result = new List<IList<int>>();
        List<int> subset = new List<int>();
        GenerateSubsets(nums, 0, subset, result);
        return result;
    }

    private void GenerateSubsets(int[] nums, int index, List<int> subset, IList<IList<int>> result) {
        result.Add(new List<int>(subset));

        for (int i = index; i < nums.Length; i++) {
            subset.Add(nums[i]);
            GenerateSubsets(nums, i + 1, subset, result);
            subset.RemoveAt(subset.Count - 1);
        }
    }
}
