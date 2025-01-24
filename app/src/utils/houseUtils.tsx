export const getHouseBackground = (house: string | undefined) => {
  if (!house?.trim()) return '/house_bg/Hogwarts.png';

  const houseImages: Record<string, string> = {
    Gryffindor: '/house_bg/Gryffindor.jpg',
    Slytherin: '/house_bg/Slytherin.jpg',
    Ravenclaw: '/house_bg/Ravenclaw.jpg',
    Hufflepuff: '/house_bg/Hufflepuff.jpg'
  };

  return houseImages[house] || '/house_bg/Hogwarts.png';
};

export const getHouseLogo = (house: string | undefined) => {
  if (!house?.trim()) return '/house_logo/Hogwarts.png';

  const houseLogos: Record<string, string> = {
    Gryffindor: '/house_logo/Gryffindor.png',
    Slytherin: '/house_logo/Slytherin.png', 
    Ravenclaw: '/house_logo/Ravenclaw.png',
    Hufflepuff: '/house_logo/Hufflepuff.png'
  };

  return houseLogos[house.trim()] || '/house_logo/Hogwarts.png';
};