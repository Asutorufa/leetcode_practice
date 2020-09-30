#include <iostream>
#include <vector>
#include <queue>

/**
给定一个二叉搜索树, 找到该树中两个指定节点的最近公共祖先。
百度百科中最近公共祖先的定义为：
    “对于有根树 T 的两个结点 p、q，
    最近公共祖先表示为一个结点 x，
    满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”
例如，给定如下二叉搜索树:  root = [6,2,8,0,4,7,9,null,null,3,5]
![](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/12/14/binarysearchtree_improved.png)

示例 1:
输入: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8
输出: 6 
解释: 节点 2 和节点 8 的最近公共祖先是 6。

示例 2:
输入: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 4
输出: 2
解释: 节点 2 和节点 4 的最近公共祖先是 2, 因为根据定义最近公共祖先节点可以为节点本身。

说明:
    所有节点的值都是唯一的。
    p、q 为不同节点且均存在于给定的二叉搜索树中。 
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

// 利用二叉搜索树的性质 很快
// 由于还没有go提交的模板 就先用c++了
class Solution
{
public:
    TreeNode *lowestCommonAncestor(TreeNode *root, TreeNode *p, TreeNode *q)
    {
        while (root != NULL)
        {
            if (root->val < p->val && root->val < q->val)
            {
                root = root->right;
            }
            else if (root->val > p->val && root->val > q->val)
            {
                root = root->left;
            }
            else
            {
                break;
            }
        }
        return root;
    }
    TreeNode *lowestCommonAncestor2(TreeNode *root, TreeNode *p, TreeNode *q)
    {
        if (root->val < p->val && root->val < q->val)
        {
            return lowestCommonAncestor2(root->right, p, q);
        }
        else if (root->val > p->val && root->val > q->val)
        {
            return lowestCommonAncestor2(root->left, p, q);
        }
        return root;
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
    }
};

int main()
{
    std::vector<int> cc = {6, 2, 8, 1, 4, 7, 9};
    TreeNode *xx = CreateTree::create(cc, 0);
    // CreateTree::level_print(xx);
    Solution s = Solution();
    TreeNode *res = s.lowestCommonAncestor(xx, new TreeNode(7), new TreeNode(9));
    CreateTree::level_print(res);
    std::cout << std::endl;
    res = s.lowestCommonAncestor2(xx, new TreeNode(7), new TreeNode(9));
    CreateTree::level_print(res);

    std::cout << std::endl
              << std::endl
              << "Tree Test" << std::endl;
    std::vector<int> c = {1, 2, 3, 4, 5};
    TreeNode *x = CreateTree::create(c, 0);
    CreateTree::print_tree(x);
    std::cout << std::endl;
    CreateTree::level_print(x);
    std::cout << std::endl;

    return 0;
}