# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Codec:

    def serialize(self, root):
        """Encodes a tree to a single string.
        
        :type root: TreeNode
        :rtype: str
        """
        if root is not None:
            this_res = str(root.val)
            left_res = self.serialize(root.left)
            right_res = self.serialize(root.right)
            res = this_res + "," + left_res + "," + right_res
        else:
            return "null"
        return res
        

    def deserialize(self, data):
        """Decodes your encoded data to tree.
        
        :type data: str
        :rtype: TreeNode
        """
        nodes = data.split(",")
        if nodes[0] == "null":
            return None
        else:
            this_node, _ = self.gen_tree(nodes)
            return this_node


    def gen_tree(self, nodes):
        if nodes[0] == "null":
            return None, 1
        this_node = TreeNode(int(nodes[0]))
        this_node.left, left_len = self.gen_tree(nodes[1:])
        this_node.right, right_len = self.gen_tree(nodes[1+left_len:])
        return this_node, 1 + left_len + right_len
        

# Your Codec object will be instantiated and called as such:
# codec = Codec()
# codec.deserialize(codec.serialize(root))