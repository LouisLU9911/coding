class Solution {
public:
    int lengthOfLongestSubstring(string s) {
        map<char, int> charMap;
        int left = 0, right = -1, maxLength = 0;
        for (int i = 0; i < s.size(); i++) {
            if (charMap.find(s[i]) != charMap.end() && charMap[s[i]] >= left) {
                // if key exists and the index of it is bigger than left, we need to update it
                left = charMap[s[i]] + 1;
            }
            charMap[s[i]] = i;
            right++;
            maxLength = max(maxLength, right - left + 1);
        }
        return maxLength;
    }

private:
    int max(int a, int b) {
        if (a > b) {
            return a;
        }
        return b;
    }
};
