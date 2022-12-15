import {MyTwoWayNode} from '../src/node';

describe('MyTwoWayNode', () => {
  it('creates a new number node', () => {
    const node3 = new MyTwoWayNode(3);

    expect(node3.getValue()).toBe(3);
    expect(node3.getNext()).toBe(null);
    expect(node3.getPrev()).toBe(null);
  })
  it('links two nodes', () => {
    const node42 = new MyTwoWayNode(42);
    const node3 = new MyTwoWayNode(3);

    node42.setNext(node3)
    node3.setPrev(node42)

    expect(node3.getNext()).toBe(null)
    expect(node42.getPrev()).toBe(null)
    expect(node42.getNext()).toBe(node3)
    expect(node3.getPrev()).toBe(node42)
  })
})