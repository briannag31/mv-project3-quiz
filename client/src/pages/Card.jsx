const CardStyle = {
    border: "1px solid black",
    padding: "20px",
    margin: "20px",
    width: "200px",
    height: "300px"
  };
const Card = ({idvCard})=>{
  const [isFlipped, setIsFlipped] = useState(false)
  console.log(idvCard)
  return(

    <ReactCardFlip isFlipped={isFlipped} flipDirection="horizontal">
      <div  style={CardStyle} onMouseEnter={()=> setIsFlipped((prev)=> !prev)} className="CardFront">
        <div>
          This is the fromt of the card
          <h3>{idvCard.term}</h3>
        </div>
      </div>
      <div style={CardStyle}
        onMouseLeave={() => setIsFlipped((prev) => !prev)}
        className="CardBack">
          This is the back of the card
          <h3>{idvCard.answer}</h3>
      </div>
    </ReactCardFlip>
  )
}

export default Card