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
      myList.setNewHead(7)

      expect(myList.stringifyList()).toBe("7-35")
    })
  })
})