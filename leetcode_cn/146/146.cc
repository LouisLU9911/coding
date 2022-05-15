#include <iostream>
#include <unordered_map>
using namespace std;

class DLinkedNode {
public:
  int key, value;
  DLinkedNode *prev;
  DLinkedNode *next;
  DLinkedNode() : key(0), value(0), prev(nullptr), next(nullptr) {}
  DLinkedNode(int _key, int _value)
      : key(_key), value(_value), prev(nullptr), next(nullptr) {}
  DLinkedNode(int _key, int _value, DLinkedNode *_prev, DLinkedNode *_next)
      : key(_key), value(_value), prev(_prev), next(_next) {}
  void show() {
    cout << "<" << key << "," << value << ">,";
    if (next) {
      next->show();
    }
  }
};

class LRUCache {
private:
  unordered_map<int, DLinkedNode *> cache;
  int size;
  int cap;
  DLinkedNode *head;
  DLinkedNode *tail;

public:
  LRUCache(int capacity) {
    head = new DLinkedNode();
    tail = new DLinkedNode();
    head->next = tail;
    tail->prev = head;
    size = 0;
    cap = capacity;
  }

  ~LRUCache() {
    delete head;
    delete tail;
  }

  void show() {
    head->show();
    cout << endl;
  }

  void update(int key, int value) {
    DLinkedNode *node = cache[key];
    node->value = value;

    node->prev->next = node->next;
    node->next->prev = node->prev;

    node->next = head->next;
    node->prev = head;
    head->next->prev = node;
    head->next = node;
  }

  void update(int key) { update(key, cache[key]->value); }

  void add2Head(int key, int value) {
    DLinkedNode *node = new DLinkedNode(key, value, head, head->next);
    cache[key] = node;
    head->next->prev = node;
    head->next = node;
    size++;
  }

  void rmTail() {
    if (size == 0) {
      return;
    }
    size--;
    DLinkedNode *node = tail->prev;
    node->prev->next = tail;
    tail->prev = node->prev;
    cache.erase(node->key);
    delete node;
  }

  int get(int key) {
    if (!cache.count(key)) {
      return -1;
    }
    update(key);
    // show();
    return cache[key]->value;
  }

  void put(int key, int value) {
    if (cache.count(key)) {
      update(key, value);
    } else {
      add2Head(key, value);
      if (size > cap) {
        rmTail();
      }
    }
    // show();
  }
};

/**
 * Your LRUCache object will be instantiated and called as such:
 * LRUCache* obj = new LRUCache(capacity);
 * int param_1 = obj->get(key);
 * obj->put(key,value);
 */
