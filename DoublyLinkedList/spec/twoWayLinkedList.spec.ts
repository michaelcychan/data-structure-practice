import {TwoWayLinkedList} from '../src/twoWayLinkedList';

describe('TwoWayLinkedList', () => {
  describe('initialises', () => {
    it('initialises with no value', () => {
      const myList = new TwoWayLinkedList();

      expect(myList.getHead()).toBe(null)
      expect(myList.getTail()).toBe(null)
    })
    it('initialises with a given value', () => {
      const myList = new TwoWayLinkedList(72);

      expect(myList.getHead()).not.toBe(null)
      expect(myList.getTail()).not.toBe(null)
      expect(myList.getHead()!.getValue()).toBe(72)
      expect(myList.getTail()!.getValue()).toBe(72)
    })
  })
  describe('length', () => {
    it('returns 0 for empty list', () => {
      const myList = new TwoWayLinkedList();
      expect(myList.length()).toBe(0)
    })
    it('returns 0 for empty list', () => {
      const myList = new TwoWayLinkedList(72);
      expect(myList.length()).toBe(1)
    })
  })
  describe('stringify', () => {
    it('returns empty string for empty list', () => {
      const myList = new TwoWayLinkedList();
      expect(myList.stringify()).toBe("")
    })
    it('returns a string for list with one node', () => {
      const myList = new TwoWayLinkedList(45);
      expect(myList.stringify()).toBe("45")
    })
    it('returns a string for list with multiple nodes', () => {
      const myList = new TwoWayLinkedList(45);
      myList.addHead(37)
      myList.addHead(24)
      myList.addHead(15)
      myList.addHead(2)
      expect(myList.stringify()).toBe("2-15-24-37-45")
    })
  })
  describe('addHead', ()=> {
    it('adds a head to an empty list', () => {
      const myList = new TwoWayLinkedList();
      
      myList.addHead(5)
      expect(myList.stringify()).toBe("5")
      expect(myList.getHead()!.getValue()).toBe(5)
      expect(myList.getTail()!.getValue()).toBe(5)
    })
    it('adds a head to a list with one element', () => {
      const myList = new TwoWayLinkedList(67);
      
      myList.addHead(5)
      expect(myList.stringify()).toBe("5-67")
      expect(myList.getHead()!.getValue()).toBe(5)
      expect(myList.getTail()!.getValue()).toBe(67)
    })
  })
  describe('addTail', () => {
    it('adds a tail to an empty list', () => {
      const myList = new TwoWayLinkedList();
      
      myList.addTail(95)
      expect(myList.stringify()).toBe("95")
      expect(myList.getHead()!.getValue()).toBe(95)
      expect(myList.getTail()!.getValue()).toBe(95)
    })
    it('adds a tail to a list with one element', () => {
      const myList = new TwoWayLinkedList(34);
      
      myList.addTail(95)
      expect(myList.stringify()).toBe("34-95")
      expect(myList.getHead()!.getValue()).toBe(34)
      expect(myList.getTail()!.getValue()).toBe(95)
    })
  })
})