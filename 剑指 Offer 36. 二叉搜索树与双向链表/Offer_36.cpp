#include <iostream>
#include <vector>
#include <queue>

/**
输入一棵二叉搜索树，将该二叉搜索树转换成一个排序的循环双向链表。要求不能创建任何新的节点，只能调整树中节点指针的指向。



为了让您更好地理解问题，以下面的二叉搜索树为例：
![](https://assets.leetcode.com/uploads/2018/10/12/bstdlloriginalbst.png)

我们希望将这个二叉搜索树转化为双向循环链表。
链表中的每个节点都有一个前驱和后继指针。
对于双向循环链表，第一个节点的前驱是最后一个节点，最后一个节点的后继是第一个节点。

下图展示了上面的二叉搜索树转化成的链表。“head” 表示指向链表中有最小元素的节点。
![](https://assets.leetcode.com/uploads/2018/10/12/bstdllreturndll.png)

特别地，我们希望可以就地完成转换操作。
当转化完成以后，树中节点的左指针需要指向前驱，树中节点的右指针需要指向后继。
还需要返回链表中的第一个节点的指针。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/er-cha-sou-suo-shu-yu-shuang-xiang-lian-biao-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/*
// Definition for a Node.
class Node {
public:
    int val;
    Node* left;
    Node* right;

    Node() {}

    Node(int _val) {
        val = _val;
        left = NULL;
        right = NULL;
    }

    Node(int _val, Node* _left, Node* _right) {
        val = _val;
        left = _left;
        right = _right;
    }
};
*/

class Node
{
public:
    int val;
    Node *left;
    Node *right;
    Node() {}
    Node(int _val)
    {
        val = _val;
        left = NULL;
        right = NULL;
    }

    Node(int _val, Node *_left, Node *_right)
    {
        val = _val;
        left = _left;
        right = _right;
    }

    static void print_link(Node *root)
    {
        while (root != NULL)
        {
            std::cout << root->val << " -> ";
            root = root->right;
        }
        std::cout << std::endl;
    }
    static void print_link2(Node *root)
    {
        while (root != NULL)
        {
            std::cout << root->val << " -> ";
            root = root->left;
        }
        std::cout << std::endl;
    }
    static void print_loop_link(Node *root)
    {
        if (root == NULL)
            return;
        int start = root->val;
        std::cout << root->val << " -> ";
        root = root->right;
        while (true)
        {
            if (root->val == start)
                break;
            std::cout << root->val << " -> ";
            root = root->right;
        }
        std::cout << "loop -> ";

        start = root->val;
        std::cout << root->val << " -> ";
        root = root->left;
        while (true)
        {
            if (root->val == start)
                break;
            std::cout << root->val << " -> ";
            root = root->left;
        }
        std::cout << std::endl;
    }
};

class Solution
{
public:
    Node *pre, *head;
    Node *treeToDoublyList(Node *root)
    {
        if (root == NULL)
            return NULL;
        pre = NULL, head = NULL; // make pre and head is NULL at start
        trans(root);
        head->left = pre;
        pre->right = head;
        return head;
    }

    void trans(Node *root)
    {
        if (root == NULL)
        {
            return;
        }
        trans(root->left);
        if (pre != NULL)
            pre->right = root;
        else
            head = root;
        root->left = pre;
        pre = root;
        trans(root->right);
    }
};

class CreateTree
{
public:
    static Node *create(std::vector<int> a, int k)
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
        Node *root = new Node(a[k]);
        root->left = create(a, k * 2 + 1);
        root->right = create(a, k * 2 + 2);
        return root;
    }

    static void print_tree(Node *root)
    {
        if (root == NULL)
        {
            return;
        }
        std::cout << root->val << " -> ";
        print_tree(root->left);
        print_tree(root->right);
    }

    static void pre_print(Node *root)
    {
        if (root == NULL)
        {
            return;
        }
        pre_print(root->left);
        std::cout << root->val << " -> ";
        pre_print(root->right);
    }

    static void level_print(Node *root)
    {
        std::queue<Node *> qu;
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
                Node *tmp = qu.front();
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
    std::vector<int> x = {4, 2, 5, 1, 3};
    Solution s = Solution();
    s.trans(CreateTree::create(x, 0));
    Node::print_link(s.head);
    Node::print_link2(s.pre);

    Node::print_loop_link(s.treeToDoublyList(CreateTree::create(x, 0)));
    std::vector<int> x2 = {4, 6, 5, 1, 3};
    Node::print_loop_link(s.treeToDoublyList(CreateTree::create(x2, 0)));
    return 0;
}