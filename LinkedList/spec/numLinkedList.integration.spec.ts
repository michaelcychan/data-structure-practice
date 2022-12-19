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

      const removed = myList.removeNode(4)
      expect(removed!).toBe(4)
      expect(myList.stringifyList()).toBe("17-32-55")
    })
    it('removes an existing target node (second node)', () => {
      const myList = new MyNumLinkedList(55)
      myList.addNewHead(32)
      myList.addNewHead(17)
      myList.addNewHead(4)
      expect(myList.stringifyList()).toBe("4-17-32-55")

      const removed = myList.removeNode(17)
      expect(removed!).toBe(17)
      expect(myList.stringifyList()).toBe("4-32-55")
    })
    it('removes an existing target node (third node)', () => {
      const myList = new MyNumLinkedList(55)
      myList.addNewHead(32)
      myList.addNewHead(17)
      myList.addNewHead(4)
      expect(myList.stringifyList()).toBe("4-17-32-55")

      const removed = myList.removeNode(32)
      expect(removed!).toBe(32)
      expect(myList.stringifyList()).toBe("4-17-55")
    })
    it('removes an existing target node (last node)', () => {
      const myList = new MyNumLinkedList(55)
      myList.addNewHead(32)
      myList.addNewHead(17)
      myList.addNewHead(4)
      expect(myList.stringifyList()).toBe("4-17-32-55")

      const removed = myList.removeNode(55)
      expect(removed!).toBe(55)
      expect(myList.stringifyList()).toBe("4-17-32")
    })
    it('removes the only node in a list', () => {
      const myList = new MyNumLinkedList(55)
      expect(myList.stringifyList()).toBe("55")

      const removed = myList.removeNode(55)
      expect(removed!).toBe(55)
      expect(myList.stringifyList()).toBe("")
    })
    it('removes an non-existing node', () => {
      const myList = new MyNumLinkedList(55)
      myList.addNewHead(32)
      myList.addNewHead(17)
      myList.addNewHead(4)
      expect(myList.stringifyList()).toBe("4-17-32-55")

      jest.spyOn(global.console, 'error')
      const removed = myList.removeNode(1)
      expect(removed).toBe(null)
      expect(console.error).toBeCalled()
      expect(myList.stringifyList()).toBe("4-17-32-55")
    })
    it('removes an non-existing from an empty list', () => {
      const myList = new MyNumLinkedList()
      expect(myList.stringifyList()).toBe("")

      jest.spyOn(global.console, 'error')
      const removed = myList.removeNode(1)
      expect(removed).toBe(null)
      expect(console.error).toBeCalled()
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
  describe("swapping", () => {
    it('swaps two exisitng nodes', () => {
      const myList = new MyNumLinkedList(55)
      myList.addNewHead(32)
      myList.addNewHead(17)
      myList.addNewHead(4)
      expect(myList.stringifyList()).toBe("4-17-32-55")

      myList.swapping(17, 32)
      expect(myList.stringifyList()).toBe("4-32-17-55")
    })
    it('swaps two exisitng nodes including tail', () => {
      const myList = new MyNumLinkedList(55)
      myList.addNewHead(32)
      myList.addNewHead(17)
      myList.addNewHead(4)
      expect(myList.stringifyList()).toBe("4-17-32-55")

      myList.swapping(17, 55)
      expect(myList.stringifyList()).toBe("4-55-32-17")
    })
    it('swaps two exisitng nodes including head', () => {
      const myList = new MyNumLinkedList(55)
      myList.addNewHead(32)
      myList.addNewHead(17)
      myList.addNewHead(4)
      expect(myList.stringifyList()).toBe("4-17-32-55")

      myList.swapping(4, 32)
      expect(myList.stringifyList()).toBe("32-17-4-55")
    })
    it('swaps two exisitng nodes including head (case 2)', () => {
      const myList = new MyNumLinkedList(55)
      myList.addNewHead(32)
      myList.addNewHead(17)
      myList.addNewHead(4)
      expect(myList.stringifyList()).toBe("4-17-32-55")

      myList.swapping(32, 4)
      expect(myList.stringifyList()).toBe("32-17-4-55")
    })
    it('swaps with one non-existing node', () => {
      const myList = new MyNumLinkedList(55)
      myList.addNewHead(32)
      myList.addNewHead(17)
      myList.addNewHead(4)
      expect(myList.stringifyList()).toBe("4-17-32-55")

      jest.spyOn(global.console, 'error')
      
      myList.swapping(33, 4)
      expect(console.error).toBeCalled()
      expect(myList.stringifyList()).toBe("4-17-32-55")
    })
  })
  describe('findMid', () => {
    it('returns null for empty list', () => {
      const myList = new MyNumLinkedList()

      expect(myList.findMid()).toBe(null)
    })
    it('returns head for single node list', () => {
      const myList = new MyNumLinkedList(55)
      expect(myList.findMid()).toBe(55)
    })
    it('returns middle for odd number nodes', () => {
      const myList = new MyNumLinkedList(55)
      myList.addNewHead(32)
      myList.addNewHead(5)

      expect(myList.findMid()).toBe(32)
    })
    it('returns middle node to the tail for even number nodes', () => {
      const myList = new MyNumLinkedList(55)
      myList.addNewHead(32)
      myList.addNewHead(23)
      myList.addNewHead(5)

      expect(myList.findMid()).toBe(32)
    })
  })
  describe('length', () => {
    it('returns 0 for empty list', () => {
      const myList = new MyNumLinkedList()
      expect(myList.length()).toBe(0)
    })
    it('returns 1 for a single node list', () => {
      const myList = new MyNumLinkedList(47)
      expect(myList.length()).toBe(1)
    })
    it('returns 1 for a single node list', () => {
      const myList = new MyNumLinkedList(55)
      myList.addNewHead(32)
      myList.addNewHead(23)
      myList.addNewHead(5)
      expect(myList.length()).toBe(4)
    })
  })
})