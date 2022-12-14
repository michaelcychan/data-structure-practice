import {MyNumberNode} from './node'
import {NodeType} from './node'

type initValueType = number | null

export class MyNumLinkedList {
  #head: NodeType

  constructor(value:initValueType = null) {
    if (value != null) {
      this.#head = new MyNumberNode(value)
    } else {
      this.#head = null
    }
  }

  getHead() {
    return this.#head
  }

  addNewHead(newValue: number) {
    const newNode = new MyNumberNode(newValue)
    if (this.#head != null) {
      newNode.setNext(this.#head)
    }
    this.#head = newNode
  }

  stringifyList(){
    let current: NodeType = this.getHead()
    let outputString = ""
    while (current) {
      outputString += current.getValue().toString()
      const nextNode = current.getNext()
      if (nextNode) {
        outputString += "-"
      }
      current = nextNode
    }
    return outputString
  }

  setNewHead(newNode: NodeType) {
    this.#head = newNode
  }

  removeNode(value:number) {
    let current: NodeType = this.getHead()
    if (current) {
      if (current.getValue() === value) {
        this.setNewHead(current.getNext())
        return
      }
    }
    while (current) {
      const nextNode = current.getNext()

      if (nextNode) {
        if (nextNode.getValue() === value) {
          current.setNext(nextNode.getNext())
          break
        } else {
          current = current.getNext()
        }
      } else {
        console.error(`target: ${value} not found, nothing was removed`)
        break
      }
    }
  }
  
  
}