#include <iostream>
#include <string>
#include <queue>
#include <vector>
using namespace std;
/**
请实现两个函数，分别用来序列化和反序列化二叉树。
示例:

你可以将以下二叉树：
    1
   / \
  2   3
     / \
    4   5

序列化为 "[1,2,3,null,null,4,5]"
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

// 都比较简单 魔改一下层次遍历就行了..
class Codec
{
public:
    // Encodes a tree to a single string.
    string serialize(TreeNode *root)
    {
        if (root == NULL)
            return "[]";
        std::queue<TreeNode *> qu;
        qu.push(root);
        std::string s = "[";
        while (true)
        {
            int l = qu.size();
            if (l == 0)
                break;
            // std::cout << "queue size: " << l << std::endl;
            for (int i = 0; i < l; i++)
            {
                TreeNode *tmp = qu.front();
                qu.pop();
                // std::cout << tmp->val << " -> ";
                if (tmp == NULL)
                {
                    s.append("null,");
                    continue;
                }
                else
                    s.append(to_string(tmp->val) + ",");
                qu.push(tmp->left);
                qu.push(tmp->right);
            }
        }
        s[s.size() - 1] = ']';
        // std::cout << std::endl;
        return s;
    }

    // Decodes your encoded data to tree.
    TreeNode *deserialize(string data)
    {
        if (data == "[]")
            return NULL;
        // std::cout << data << std::endl;
        std::vector<std::string> tmp;
        int i = 1;
        int sub_size = 0;
        int last = 1;
        while (i < data.size())
        {
            last = i;
            sub_size = 0;
            while (data[i] != ',' && data[i] != ']')
            {
                sub_size++;
                i++;
                // std::cout << data[i] << " ";
            }
            i++;
            // std::cout << "sub->" << data.substr(last, sub_size) << std::endl;
            tmp.push_back(data.substr(last, sub_size));
        }
        // std::cout << std::endl;
        // for (int i = 0; i < tmp.size(); i++)
        // std::cout << tmp.at(i) << ' ';
        // std::cout << std::endl;
        return create(tmp);
    }

    static TreeNode *create(std::vector<std::string> a)
    {
        TreeNode *root = new TreeNode(std::stoi(a[0]));
        std::queue<TreeNode *> qu;
        qu.push(root);
        int i = 1;
        while (qu.size() != 0 && i < a.size())
        {
            TreeNode *node = qu.front();
            qu.pop();
            if (a[i] != "null")
            {
                node->left = new TreeNode(std::stoi(a[i]));
                qu.push(node->left);
            }
            if (a[i + 1] != "null")
            {
                node->right = new TreeNode(std::stoi(a[i + 1]));
                qu.push(node->right);
            }
            i += 2;
        }
        return root;
    }
};

// Your Codec object will be instantiated and called as such:
// Codec codec;
// codec.deserialize(codec.serialize(root));

class CreateTree
{
public:
    static TreeNode *create(std::vector<int> a, int k)
    {
        if (a.size() == 0 || a.size() - 1 < k)
            return NULL;
        if (a[k] == 0)
            return NULL;
        // std::cout << a[k] << std::endl;
        TreeNode *root = new TreeNode(a[k]);
        root->left = create(a, k * 2 + 1);
        root->right = create(a, k * 2 + 2);
        return root;
    }

    static TreeNode *create(std::vector<std::string> a)
    {
        TreeNode *root = new TreeNode(std::stoi(a[0]));
        std::queue<TreeNode *> qu;
        qu.push(root);
        int i = 1;
        while (qu.size() != 0 && i < a.size())
        {
            TreeNode *node = qu.front();
            qu.pop();
            if (a[i] != "null")
            {
                node->left = new TreeNode(std::stoi(a[i]));
                qu.push(node->left);
            }
            if (a[i + 1] != "null")
            {
                node->right = new TreeNode(std::stoi(a[i + 1]));
                qu.push(node->right);
            }
            i += 2;
        }
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

    static void pre_print(TreeNode *root)
    {
        if (root == NULL)
        {
            return;
        }
        pre_print(root->left);
        std::cout << root->val << " -> ";
        pre_print(root->right);
    }

    static void level_print(TreeNode *root)
    {
        if (!root)
            return;
        std::queue<TreeNode *> qu;
        qu.push(root);
        while (1)
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
    std::vector<int> a = {4, 2, 5, 1, 3};
    Codec c = Codec();
    std::cout << c.serialize(CreateTree::create(a, 0)) << endl;
    a = {1, 2, 3, 0, 0, 4, 5};
    std::cout << c.serialize(CreateTree::create(a, 0)) << endl;
    CreateTree::level_print(c.deserialize("[1,2,3,null,null,4,5]"));
    CreateTree::level_print(c.deserialize("[5,2,3,null,null,2,4,3,1]"));
    CreateTree::level_print(c.deserialize("[]"));
    return 0;
}