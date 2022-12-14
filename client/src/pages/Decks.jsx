export default function Decks({cards, setCards, setAddCard,card,setCard,newCard, setNewCard, addCard, setViewCards, setViewDecks, decks, setDecks, deck, setDeck}){
    const url = "http://localhost:3001/api"
    const changeView = async()=>{
        setViewDecks(false)
        setViewCards(true)
        const res = await fetch(`${url}/decks/639775f29df0ccec2b1428fc`)

      const data = await res.json()
      console.log("res is: ",data)
      setDecks(data)
    }
    console.log("decks are:",decks)
    return (  <>
        <button onClick={()=>setViewDecks(false)}>Go Back Home</button>
        {decks.map((idvDeck, deckIdx)=>{
            return(
                <div className="plainCard" key={deckIdx}>
                    <h3 onClick={()=>{changeView()}} >{idvDeck.deckname}</h3>
                </div>
            )
        })}
        {/* {cards.map((idvCard, idx)=>{
            return <h3>{idvCard.term}</h3>
        })} */}
  
         </>)
    
}