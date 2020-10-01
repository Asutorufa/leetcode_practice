#include <iostream>
#include <vector>
#include <queue>
/**
 * 给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。
百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，
最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”
例如，给定如下二叉树:  root = [3,5,1,6,2,0,8,null,null,7,4]
![](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/12/15/binarytree.png)
示例 1:
输入: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
输出: 3
解释: 节点 5 和节点 1 的最近公共祖先是节点 3。

示例 2:
输入: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
输出: 5
解释: 节点 5 和节点 4 的最近公共祖先是节点 5。因为根据定义最近公共祖先节点可以为节点本身。

说明:
    所有节点的值都是唯一的。
    p、q 为不同节点且均存在于给定的二叉树中。
*/

/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode(int x) : val(x), left(NULL), right(NULL) {}
 * };
 */

struct TreeNode
{
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode(int x) : val(x), left(NULL), right(NULL) {}
};

class Solution
{
public:
    TreeNode *lowestCommonAncestor(TreeNode *root, TreeNode *p, TreeNode *q)
    {
        if (root == NULL || root->val == p->val || root->val == q->val)
            return root;                                           // 当root为NULL或者是查询的某一节点时, 返回root
        TreeNode *left = lowestCommonAncestor(root->left, p, q);   // 递归
        TreeNode *right = lowestCommonAncestor(root->right, p, q); // 递归
        if (left == NULL)
            return right;
        if (right == NULL) // 当right为NULL时, 返回left, 当然left也有可能为NULL
            return left;
        return root; // 当 right 和 left 都不为 NULL时, 说明当前节点为最短根节点, 返回root
    }
};

class CreateTree
{
public:
    static TreeNode *create(std::vector<int> a, int k)
    {
        if (a.size() == 0 || a.size() - 1 < k)
        {
            return NULL;
        }
        if (a[k] == 0)
        {
            return NULL;
        }
        // std::cout << a[k] << std::endl;
        TreeNode *root = new TreeNode(a[k]);
        root->left = create(a, k * 2 + 1);
        root->right = create(a, k * 2 + 2);
        return root;
    }
    static void print_tree(TreeNode *root)
    {
        if (root == NULL)
        {
            return;
        }
        std::cout << root->val << " -> ";
        print_tree(root->left);
        print_tree(root->right);
    }
    static void level_print(TreeNode *root)
    {
        std::queue<TreeNode *> qu;
        qu.push(root);
        while (true)
        {
            int l = qu.size();
            if (l == 0)
            {
                break;
            }
            // std::cout << "queue size: " << l << std::endl;
            for (int i = 0; i < l; i++)
            {
                TreeNode *tmp = qu.front();
                qu.pop();
                std::cout << tmp->val << " -> ";
                if (tmp->left != NULL)
                {
                    qu.push(tmp->left);
                }
                if (tmp->right != NULL)
                {
                    qu.push(tmp->right);
                }
            }
        }
        std::cout << std::endl;
    }
};

int main()
{
    std::vector<int> a = {3, 5, 1, 6, 2, 9, 8, 0, 0, 7, 4};
    TreeNode *x = CreateTree::create(a, 0);
    // CreateTree::level_print(x);
    Solution s = Solution();
    CreateTree::level_print(s.lowestCommonAncestor(x, new TreeNode(5), new TreeNode(1))); // start at 3
    CreateTree::level_print(s.lowestCommonAncestor(x, new TreeNode(5), new TreeNode(4))); // start at 5
    return 0;
}