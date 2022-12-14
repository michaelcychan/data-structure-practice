import {MyNumberNode} from '../src/node';
import {MyNumLinkedList} from '../src/numLinkedList';

describe('MyNumberLinkedList', () => {
  describe('initiates', () => {
    it ("with an empty value without throwing an error", () => {
      const myList = new MyNumLinkedList();
      expect(myList.stringifyList()).toBe("")
    })
    it("with a value as head", () => {
      const myList = new MyNumLinkedList(5);
      expect(myList.stringifyList()).toBe("5")
    })
  })
  describe('adds new node', () => {
    it("adds node at beginning and can print out", () => {
      const myList = new MyNumLinkedList(35);
      myList.addNewHead(7)

      expect(myList.stringifyList()).toBe("7-35")
    })
    it("adds node to an empty list", () => {
      const myList = new MyNumLinkedList();
      myList.addNewHead(7)

      expect(myList.stringifyList()).toBe("7")
    })
  })
  describe('remove a node', () => {
    it('removes an existing target node (head)', () => {
      const myList = new MyNumLinkedList(55)
      myList.addNewHead(32)
      myList.addNewHead(17)
      myList.addNewHead(4)
      expect(myList.stringifyList()).toBe("4-17-32-55")

      myList.removeNode(4)
      expect(myList.stringifyList()).toBe("17-32-55")
    })
    it('removes an existing target node (second node)', () => {
      const myList = new MyNumLinkedList(55)
      myList.addNewHead(32)
      myList.addNewHead(17)
      myList.addNewHead(4)
      expect(myList.stringifyList()).toBe("4-17-32-55")

      myList.removeNode(17)
      expect(myList.stringifyList()).toBe("4-32-55")
    })
    it('removes an existing target node (third node)', () => {
      const myList = new MyNumLinkedList(55)
      myList.addNewHead(32)
      myList.addNewHead(17)
      myList.addNewHead(4)
      expect(myList.stringifyList()).toBe("4-17-32-55")

      myList.removeNode(32)
      expect(myList.stringifyList()).toBe("4-17-55")
    })
    it('removes an existing target node (last node)', () => {
      const myList = new MyNumLinkedList(55)
      myList.addNewHead(32)
      myList.addNewHead(17)
      myList.addNewHead(4)
      expect(myList.stringifyList()).toBe("4-17-32-55")

      myList.removeNode(55)
      expect(myList.stringifyList()).toBe("4-17-32")
    })
    it('removes the only node in a list', () => {
      const myList = new MyNumLinkedList(55)
      expect(myList.stringifyList()).toBe("55")

      myList.removeNode(55)
      expect(myList.stringifyList()).toBe("")
    })
    it('removes an non-existing node', () => {
      const myList = new MyNumLinkedList(55)
      myList.addNewHead(32)
      myList.addNewHead(17)
      myList.addNewHead(4)
      expect(myList.stringifyList()).toBe("4-17-32-55")

      myList.removeNode(1)
      expect(myList.stringifyList()).toBe("4-17-32-55")
    })
    it('removes an non-existing from an empty list', () => {
      const myList = new MyNumLinkedList()
      expect(myList.stringifyList()).toBe("")

      myList.removeNode(1)
      expect(myList.stringifyList()).toBe("")
    })
  })
  describe('reverseIndex', () => {
    it('returns value counting from end of tail', () => {
      const myList = new MyNumLinkedList(55)
      myList.addNewHead(32)
      myList.addNewHead(17)
      myList.addNewHead(4)
      expect(myList.stringifyList()).toBe("4-17-32-55")

      expect(myList.reverseIndex(2)).toBe(32)
    })
    it('returns value counting from end of tail', () => {
      const myList = new MyNumLinkedList(55)
      myList.addNewHead(32)
      myList.addNewHead(17)
      myList.addNewHead(4)
      expect(myList.stringifyList()).toBe("4-17-32-55")

      expect(myList.reverseIndex(1)).toBe(55)
    })
    it('returns null if length of Linked List shorter than reverse index', () => {
      const myList = new MyNumLinkedList(55)
      myList.addNewHead(32)
      myList.addNewHead(17)
      myList.addNewHead(4)
      expect(myList.stringifyList()).toBe("4-17-32-55")
      expect(myList.reverseIndex(5)).toBe(null)
    })
    it('returns first node if length of Linked List equals reverse index', () => {
      const myList = new MyNumLinkedList(55)
      myList.addNewHead(32)
      myList.addNewHead(17)
      myList.addNewHead(4)
      expect(myList.stringifyList()).toBe("4-17-32-55")

      expect(myList.reverseIndex(4)).toBe(4)
    })
    it('returns null for empty list', () => {
      const myList = new MyNumLinkedList()
      expect(myList.stringifyList()).toBe("")

      expect(myList.reverseIndex(4)).toBe(null)
    })
  })
})