import React, {useState} from "react";
import ReactCardFlip from 'react-card-flip';
// export default function Index({setViewCards,cards, setCards, setAddCard,card,setCard,newCard, setNewCard, addCard, fetchCards}){



const Card = ({project})=>{
  const [isFlipped, setIsFlipped] = useState(false)
  
  return(
    <ReactCardFlip isFlipped={isFlipped} flipDirection="horizontal" >
     
        <div   className="CardFront" onClick={()=> {setIsFlipped((prev)=> !prev)}}>
        <button type="button" class="btn-close  pull-right close-button" aria-label="Close"></button>
          <h3 className="cardFont">{project.term}</h3>
        </div>
      <div className="CardBack   align-column" >
        
          <h3 className="cardFont">{project.answer}</h3>
          <button type="button" class="btn-close  pull-right close-button-back" aria-label="Close"></button>
          <div style={{ margin: "0 auto"}} onClick={() => setIsFlipped((prev) => !prev)} >
           <p style={{marginTop:" 120px", marginBottom: "-20px"}}> Did you know the answer?</p> <br />
           <div className="card-footer">
            
            <div className="row">
              <div className="col-6 pr-0">
              <button type="button" className="btn btn-success"> <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-check2" viewBox="0 0 16 16">
  <path d="M13.854 3.646a.5.5 0 0 1 0 .708l-7 7a.5.5 0 0 1-.708 0l-3.5-3.5a.5.5 0 1 1 .708-.708L6.5 10.293l6.646-6.647a.5.5 0 0 1 .708 0z"/>
</svg>&nbsp;&nbsp;&nbsp;Yes</button>
              </div>
              <div className="col-6 pr-0">
              <button type="button" className="btn btn-danger"> <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-x-lg" viewBox="0 0 16 16">
  <path d="M2.146 2.854a.5.5 0 1 1 .708-.708L8 7.293l5.146-5.147a.5.5 0 0 1 .708.708L8.707 8l5.147 5.146a.5.5 0 0 1-.708.708L8 8.707l-5.146 5.147a.5.5 0 0 1-.708-.708L7.293 8 2.146 2.854Z"/>
</svg>&nbsp;&nbsp;&nbsp; No</button>
              </div>
            </div>
           </div>
            
            
          </div>
</div>
    </ReactCardFlip>
  )
}

const Index = ({setViewCards,cards, setCards, setAddCard,card,setCard,newCard, setNewCard, addCard, fetchCards}) =>{
  const url = "http://localhost:3001/api"

    const deleteCard = async (id) =>{
        const res = await fetch(`${url}/deleteCards/${id}`, {
          method: "DELETE",
          header: {
            "Content-Type": "application/x-www-form-urlencoded"  
          },
        })
        const data = await res.json()
        fetchCards()
      }
  console.log(cards)
   let color
  return(
   
    <div className="row hello" >
      {cards.map((item, idx)=>{
      return  item.known ? 
          
          <div className="plainCard col-3" style={{borderColor: "green", margin: "40px", marginTop: "100px"}} >
          <Card project={item} key={idx} />
         </div>
         :
         <div className="plainCard col-3" style={{borderColor: "red", margin: "40px", marginTop: "100px"}} >
         <Card project={item} key={idx} />
        </div>
       
        
        
      })}
    </div>
  )
}

export default Index
//   const [flip, setFlip] = useState(false) 
//   const url = "http://localhost:3001/api"

//     const deleteCard = async (id) =>{
//         const res = await fetch(`${url}/deleteCards/${id}`, {
//           method: "DELETE",
//           header: {
//             "Content-Type": "application/x-www-form-urlencoded"  
//           },
//         })
//         const data = await res.json()
//         fetchCards()
//       }

//       const handleClick = async(id) => {
//         const res = await fetch(`${url}/cards/${id}`)
//         const data = await res.json()
//         console.log("res is: ",res)
//         setCard(data)
//         setFlip(!flip)
//     }
//   return (  <>
//     <button onClick={()=>setAddCard(true)}>Add a New Card</button>
//     <button onClick={()=>setViewCards(false)}>Go Back Home</button>
//   <div>
//    { cards.map((individualCard, idx) =>{

//     // if (id === individualCard._id){

//     // }
//     let border = "solid red"
//     if( individualCard.known === true){
//         border = "solid green"
//     }
//     console.log(individualCard)
//       return <div style={{border: border}} className={"card"}onClick={()=>{const id = (individualCard._id);console.log("id is:",id); if (id === individualCard._id){handleClick(id)} else{(setFlip(flip))};}} >
//         {!flip ? <>
//         <h1 key={idx}>{individualCard.term}</h1>
//         <h6>{individualCard.collection}</h6>
//         {/* <h6>{individualCard.deck.deckname}</h6> */}
//       </>
//       :
//       <>
//        <h1 key={idx}>{individualCard.answer}</h1>
//        <button onClick={()=> setFlip(flip)}>Did you know the answer</button>
//       </>
//       }
        
//         <button type="button" className="btn btn-danger"  onClick={()=> { const id = (individualCard._id); deleteCard(id)}}>Delete</button>
//       </div> 
//     })}</div>
//     </>)
// }