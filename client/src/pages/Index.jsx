
export default function Index({cards, card}){
  return (  
    cards.map((individualCard, idx) =>{
      return <h1>{individualCard.term}</h1>
      console.log("individual card: ",individualCard.term)
    }))
}