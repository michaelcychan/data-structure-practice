import {MyNumberNode} from './node'
import {NodeType} from './node'

type ValueType = number | null

export class MyNumLinkedList {
  #head: NodeType

  constructor(value:ValueType = null) {
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

  length():number {
    let counter = 0
    let pointer = this.#head
    while (pointer) {
      counter += 1
      pointer = pointer.getNext()
    }
    return counter
  }

  removeNode(value:number) {
    let current: NodeType = this.getHead()
    if (current) {
      if (current.getValue() === value) {
        this.setNewHead(current.getNext())
        return current.getValue()
      }
    } else {
      console.error("List is empty, nothing is removed")
      return null
    }
    while (current) {
      const nextNode = current.getNext()

      if (nextNode) {
        if (nextNode.getValue() === value) {
          current.setNext(nextNode.getNext())
          return nextNode.getValue()
        } else {
          current = current.getNext()
        }
      
      } else {
        console.error(`target: ${value} not found, nothing was removed`)
        return null
      }
    }
  }

  reverseIndex(fromLast:number):ValueType {
    let tailPointer = this.getHead();
    let targetPointer = this.getHead();
    const head = this.getHead()
    let counter = 0;
    while (tailPointer) {
      tailPointer = tailPointer.getNext()
      counter += 1
      if (fromLast < counter ) {
        if (targetPointer) {
          targetPointer = targetPointer.getNext()
        }
      }
    }
    if (counter == fromLast && head) {
      return head.getValue()
    } else if (counter < fromLast) {
      return null
    }
    if (targetPointer) {
      return targetPointer.getValue()
    }
    return null
  }

  swapping(a:number , b: number) {
    let currentA = this.getHead()
    let currentB = this.getHead()

    let prevA: NodeType = null;
    let prevB: NodeType = null;

    while (currentA) {
      if (currentA) {
        if (currentA.getValue() === a) {
          break
        } else {
          prevA = currentA
          currentA = currentA.getNext()
        }
      }
    }

    while (currentB) {
      if (currentB) {
        if (currentB.getValue() === b) {
          break
        } else {
          prevB = currentB
          currentB = currentB.getNext()
        }
      }
    }

    if (currentA && currentB) {
      
      // cannot reverse order, will cause ciruclar link
      const tempA = currentA
      if (prevA) {
        prevA.setNext(currentB)
      } else {
        this.setNewHead(currentB)
      }
      if (prevB) {
        prevB.setNext(tempA)
      } else {
        this.setNewHead(currentA)
      }
      
      const tempNextA = currentA.getNext()
      currentA.setNext(currentB.getNext())
      currentB.setNext(tempNextA)
      
    } else {
      console.error("Error: target node not found")
    } 
  }

  findMid() {
    let rabbit = this.getHead()
    let tortoise = this.getHead()
    if (rabbit) {
      while (rabbit && tortoise) {
        const rabbitNext = rabbit.getNext()
        if (rabbitNext && tortoise) {
          rabbit = rabbitNext.getNext()
          tortoise = tortoise.getNext()
        } else {
          return tortoise.getValue()
        }
      }
      if (tortoise) {
        return tortoise.getValue()
      }
      
    } else {
      return null
    }
  }
}