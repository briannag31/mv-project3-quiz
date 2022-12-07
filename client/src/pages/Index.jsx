
export default function Index({cards, setCards, setAddCard,card,setCard,newCard, setNewCard}){
  
  return (  <>
   {/* show all cards and known cards have a green outline, unknown have red outline 
    try to use pagination or something similar to move left and right between cards? */}
    <button onClick={()=> setNewCard(true)}>Add a New Card</button>
  
   { cards.map((individualCard) =>{
    let border = "solid red"
    if( individualCard.known == true){
        border = "solid green"
    }
      return <div style={{border: border}}>
        <h1 key={individualCard._id}>{individualCard.term}</h1>
        <h6>{individualCard.collection}</h6>
      </div> 
    })}
    </>)
}