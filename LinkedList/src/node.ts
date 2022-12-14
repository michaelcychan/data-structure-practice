type nextNode = MyNumberNode | null

export class MyNumberNode {
  #value: number;
  #next: nextNode

  constructor(value:number, next:nextNode = null){
    this.#value = value;
    this.#next = next;
  }

  getValue(): number{
    return this.#value;
  }

  setValue(newValue:number) {
    this.#value = newValue;
  }

  getNext(): nextNode{
    return this.#next;
  }

  setNext(next: nextNode){
    this.#next = next;
  }
}