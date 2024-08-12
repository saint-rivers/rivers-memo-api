--
-- PostgreSQL database dump
--

-- Dumped from database version 16.3
-- Dumped by pg_dump version 16.3

-- Started on 2024-08-12 05:44:23 UTC

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- TOC entry 218 (class 1259 OID 221185)
-- Name: note_in_seq; Type: SEQUENCE; Schema: public; Owner: flashcards_owner
--

CREATE SEQUENCE public.note_in_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.note_in_seq OWNER TO flashcards_owner;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 216 (class 1259 OID 180231)
-- Name: note; Type: TABLE; Schema: public; Owner: flashcards_owner
--

CREATE TABLE public.note (
    id bigint DEFAULT nextval('public.note_in_seq'::regclass) NOT NULL,
    link character varying NOT NULL,
    title character varying,
    description character varying,
    image_url character varying
);


ALTER TABLE public.note OWNER TO flashcards_owner;

--
-- TOC entry 217 (class 1259 OID 180251)
-- Name: note_tag; Type: TABLE; Schema: public; Owner: flashcards_owner
--

CREATE TABLE public.note_tag (
    note bigint NOT NULL,
    tag character varying(24) NOT NULL
);


ALTER TABLE public.note_tag OWNER TO flashcards_owner;

--
-- TOC entry 215 (class 1259 OID 172032)
-- Name: tag; Type: TABLE; Schema: public; Owner: flashcards_owner
--

CREATE TABLE public.tag (
    tag_name character varying(24) NOT NULL
);


ALTER TABLE public.tag OWNER TO flashcards_owner;

--
-- TOC entry 3341 (class 0 OID 180231)
-- Dependencies: 216
-- Data for Name: note; Type: TABLE DATA; Schema: public; Owner: flashcards_owner
--

COPY public.note (id, link, title, description, image_url) FROM stdin;
2364	https://github.com/dyatlov/go-opengraph	GitHub - dyatlov/go-opengraph: Golang package for parsing OpenGraph data from HTML into regular structures	Golang package for parsing OpenGraph data from HTML into regular structures - dyatlov/go-opengraph	https://opengraph.githubassets.com/6a39f4aa8a862a34a2ef3ed3a5429b9aef8402f94e08f9bc1c9f756155e326ad/dyatlov/go-opengraph
2362	https://youtu.be/VmFu_j6iWKM?si=TH9M-yuvIg-_nTf4	Jon Bellion - Crop Circles (Acoustic Vertical Video)	Music video by Jon Bellion performing Crop Circles (Acoustic Vertical Video). © 2019 Capitol Records, LLChttp://vevo.ly/HnZftT	https://i.ytimg.com/vi/VmFu_j6iWKM/maxresdefault.jpg
2316	https://youtube.com/shorts/cJP455my5TY?si=hkoX8u-CV0i-YKSU	Korean is So Easy 😮 #korean	Korean is really easy with this Magic word 네!😂You can deliver so many things with this one word.    1. Yes - 네    2. I understand!- 아~네!    3. I don’t under...	https://i.ytimg.com/vi/cJP455my5TY/oardefault.jpg?sqp=-oaymwEkCJUDENAFSFqQAgHyq4qpAxMIARUAAAAAJQAAyEI9AICiQ3gB&rs=AOn4CLDjcVGXF1rCBGK3HSSXq7anmlD4kA
2358	https://youtube.com/shorts/AW6r7bhyI78?si=hnWkD5le8MtVmS73	carry me or i'm breaking up with you	he had to read chatfeaturing @itzmasayoshi FOLLOW ME    Twitch ➜ https://www.twitch.tv/QuarterJade    YouTube (2nd channel) ➜ https://youtube.com/justjodi   ...	https://i.ytimg.com/vi/AW6r7bhyI78/oar2.jpg?sqp=-oaymwEkCJUDENAFSFqQAgHyq4qpAxMIARUAAAAAJQAAyEI9AICiQ3gB&rs=AOn4CLCZwYx5NvqGxX_lzL7NEXynGnYEbw
2371	https://www.instagram.com/reel/C81_tmtq27A/?igsh=NGczZXVpZHlsY2xp	Cattos Aura on Instagram: "how does it taste??\n\nNo problem! Here's a simple and delicious recipe for homemade chicken nuggets:\n\n### Ingredients:\n- 2 boneless, skinless chicken breasts\n- 1 cup all-purpose flour\n- 1 teaspoon salt\n- 1 teaspoon black pepper\n- 1 teaspoon garlic powder\n- 1 teaspoon onion powder\n- 1 teaspoon paprika\n- 2 large eggs\n- 2 tablespoons water\n- 2 cups breadcrumbs (Panko or regular)\n- Oil for frying (vegetable, canola, or peanut oil)\n\n### Instructions:\n\n1. **Prepare the Chicken**:\n - Cut the chicken breasts into nugget-sized pieces (about 1.5 inches each).\n\n2. **Set Up Breading Stations**:\n - In the first bowl, combine the flour, salt, pepper, garlic powder, onion powder, and paprika.\n - In the second bowl, beat the eggs with water until well combined.\n - In the third bowl, place the breadcrumbs.\n\n3. **Bread the Chicken**:\n - Dredge each piece of chicken in the flour mixture, making sure to coat it evenly.\n - Dip the floured chicken pieces into the egg mixture, allowing any excess egg to drip off.\n - Coat the chicken pieces with breadcrumbs, pressing lightly to adhere the breadcrumbs to the chicken.\n\n4. **Fry the Nuggets**:\n - Heat about 2 inches of oil in a large skillet or pot to 350°F (175°C).\n - Fry the chicken nuggets in batches, making sure not to overcrowd the pan. Cook each batch for about 4-5 minutes, or until the nuggets are golden brown and cooked through.\n - Use a slotted spoon to remove the nuggets from the oil and place them on a paper towel-lined plate to drain any excess oil.\n\n5. **Serve**:\n - Serve the chicken nuggets hot with your favorite dipping sauces, such as ketchup, barbecue sauce, honey mustard, or ranch.\n\nEnjoy your homemade chicken nuggets!"	130K likes, 169 comments - cattosaura on June 30, 2024: "how does it taste??\n\nNo problem! Here's a simple and delicious recipe for homemade chicken nuggets:\n\n### Ingredients:\n- 2 boneless, skinless chicken breasts\n- 1 cup all-purpose flour\n- 1 teaspoon salt\n- 1 teaspoon black pepper\n- 1 teaspoon garlic powder\n- 1 teaspoon onion powder\n- 1 teaspoon paprika\n- 2 large eggs\n- 2 tablespoons water\n- 2 cups breadcrumbs (Panko or regular)\n- Oil for frying (vegetable, canola, or peanut oil)\n\n### Instructions:\n\n1. **Prepare the Chicken**:\n - Cut the chicken breasts into nugget-sized pieces (about 1.5 inches each).\n\n2. **Set Up Breading Stations**:\n - In the first bowl, combine the flour, salt, pepper, garlic powder, onion powder, and paprika.\n - In the second bowl, beat the eggs with water until well combined.\n - In the third bowl, place the breadcrumbs.\n\n3. **Bread the Chicken**:\n - Dredge each piece of chicken in the flour mixture, making sure to coat it evenly.\n - Dip the floured chicken pieces into the egg mixture, allowing any excess egg to drip off.\n - Coat the chicken pieces with breadcrumbs, pressing lightly to adhere the breadcrumbs to the chicken.\n\n4. **Fry the Nuggets**:\n - Heat about 2 inches of oil in a large skillet or pot to 350°F (175°C).\n - Fry the chicken nuggets in batches, making sure not to overcrowd the pan. Cook each batch for about 4-5 minutes, or until the nuggets are golden brown and cooked through.\n - Use a slotted spoon to remove the nuggets from the oil and place them on a paper towel-lined plate to drain any excess oil.\n\n5. **Serve**:\n - Serve the chicken nuggets hot with your favorite dipping sauces, such as ketchup, barbecue sauce, honey mustard, or ranch.\n\nEnjoy your homemade chicken nuggets!". 	https://scontent-ssn1-1.cdninstagram.com/v/t51.29350-15/449424499_2287808374895802_316223584083566389_n.jpg?stp=cmp1_dst-jpg_s640x640&_nc_cat=104&ccb=1-7&_nc_sid=18de74&_nc_ohc=76d7UvAaFpcQ7kNvgHtZOQc&_nc_ht=scontent-ssn1-1.cdninstagram.com&oh=00_AYDDCjq4yC3618s0oIPIq8jRN621tMJV0nhX5B2WoiAt7Q&oe=669C0E12
2315	https://youtube.com/shorts/jQXq_iWiEqE?si=9kW-IcTswVXCiWkd	Reverse Pec Deck Flyes vs Cable Rear Delt Flyes		https://i.ytimg.com/vi/jQXq_iWiEqE/oardefault.jpg?sqp=-oaymwEkCJUDENAFSFqQAgHyq4qpAxMIARUAAAAAJQAAyEI9AICiQ3gB&rs=AOn4CLAothaP46vf0BHnw_bVaUeIE6HQaw
2347	https://koreajoongangdaily.joins.com/news/2024-06-21/englishStudy/bilingualNews/Fix-the-population-crisis-before-its-too-late-KOR/2073569	Fix the population crisis before it’s too late (KOR)	The government must pay heed to advice from experts to explore measures to raise the adaptability and sustainability of society under the given population conditions instead of setting a specific population goal. 	https://koreajoongangdaily.joins.com/resources/images/common/ja_opengraph_img.png
2346	https://kids.donga.com/?ptype=article&no=20240711131532891613	큰 수박 대신 작은 수박이 대세된 이유는? 농부도, 소비자도, 유통업계도 좋아해요! 	푹푹 찌는 무더위. 냉장고에서 갓 꺼낸 차가운 수박을 한입 크게 ‘왕’ 하고 깨물어 먹는 것만큼 행복한 일도 없을 거예요. 이런 수박은 크면 클수록 맛이 좋다고 알려졌지만, 이는 ...	http://kids.donga.com/www/data/article/thum/202407/20240711131532.jpg
2345	https://www.youtube.com/watch?v=8BMXHQg9WqE&list=TLPQMTIwNzIwMjQloK0oFjgfNA&index=2	I Prevail - Bad Things (Official Music Video)	‘Bad Things’ taken from the I Prevail album ‘TRUE POWER’Stream & Download 'Bad Things' - https://found.ee/badthingsListen to ’TRUE POWER’ - https://found.ee/...	https://i.ytimg.com/vi/8BMXHQg9WqE/maxresdefault.jpg
2291	https://encore.dev/blog/queueing	Queueing – An interactive study of queueing strategies – Encore Blog	In this blog, we go on an interactive journey to understand common queueing strategies for handling HTTP requests.	https://encore.dev/assets/blog/card/queueing_cover.png
2289	https://www.instagram.com/cyuuka1010/	中華東東　（松飛台） (@cyuuka1010) • Instagram photos and videos	114K Followers, 65 Following, 826 Posts - See Instagram photos and videos from 中華東東　（松飛台） (@cyuuka1010)	https://scontent-ssn1-1.cdninstagram.com/v/t51.2885-19/137635195_161954545361435_8625592549040146814_n.jpg?stp=dst-jpg_s100x100&_nc_cat=108&ccb=1-7&_nc_sid=3fd06f&_nc_ohc=VtcHnJS8LvQQ7kNvgG9Jm4H&_nc_ht=scontent-ssn1-1.cdninstagram.com&oh=00_AYB6lAsGrUrUmDcqNkvIVnLQPxe5cpq1KHKKhufuXiNthQ&oe=669C1B46
2363	https://youtu.be/eJ-5hwr0ruY?si=he7ue6F1hoR7ZsSX	ILLENIUM, Jon Bellion - Good Things Fall Apart (Stripped / Audio)	Pre-order my upcoming ASCEND album here: https://Illenium.lnk.to/ASCENDGet tickets to see me on tour this Fall here: http://illenium.com/“Good Things Fall Ap...	\N
2343	https://youtube.com/shorts/ipRxS94nJ9E?si=kRERWwbrXXq0G2xO	Tzuyu savage moments #twice		https://i.ytimg.com/vi/ipRxS94nJ9E/oar2.jpg?sqp=-oaymwEkCJUDENAFSFqQAgHyq4qpAxMIARUAAAAAJQAAyEI9AICiQ3gB&rs=AOn4CLCxN9G4Bi5yRCVLP0jRC_5nXDR3qA
2317	https://youtube.com/shorts/NjMO51bLqF4?si=7tsSkBPubAeiHc-X	How to Order at a Korean Restaurant #learnkorean		https://i.ytimg.com/vi/NjMO51bLqF4/hq720_2.jpg?sqp=-oaymwEkCJYDENAFSFryq4qpAxYIARUAAAAAJQAAyEI9AICiQ3gB0AEB&rs=AOn4CLADf4mx2QJjXQS29gx4WkksFIDQ2w
2370	https://www.instagram.com/reel/C6XaCFQvqVq/?igsh=NjdsOG44OHQ2cGdi	꼬질꼬질 버찌🍒 on Instagram: "화난 이유: 키보드 못 밟게 해서"	6M likes, 22K comments - buzzi_ririgong on April 29, 2024: "화난 이유: 키보드 못 밟게 해서". 	https://scontent-ssn1-1.cdninstagram.com/v/t51.29350-15/441173108_7855195714569871_6560077106180884415_n.jpg?stp=cmp1_dst-jpg_s640x640&_nc_cat=1&ccb=1-7&_nc_sid=18de74&_nc_ohc=JfbQxUinm2EQ7kNvgH6bmwN&_nc_ht=scontent-ssn1-1.cdninstagram.com&oh=00_AYAImEExqcMyoR2pmgD-3PBb9w1UvuchvJtcPzkji2-3aw&oe=669C00A1
2366	https://www.youtube.com/shorts/vl_lZuDersM	@유나몽 #snsd #oh #소녀시대 #yunamong #dancechallenge #kpopdance #kpop #소울케이브 #soulcave #trending #shorts		https://i.ytimg.com/vi/vl_lZuDersM/oardefault.jpg?sqp=-oaymwEkCJUDENAFSFqQAgHyq4qpAxMIARUAAAAAJQAAyEI9AICiQ3gB&rs=AOn4CLCcbWAyaa-wET-h-XFaivADm8ouvA
2372	https://www.instagram.com/reel/C7PQqoYtdcz/?igsh=MTdwenpxb2FuY2d1cA==	Camlist • Just Pets on Instagram: "Cat’s got a purrfect balance 😂"	777K likes, 1,491 comments - getcamlist on May 21, 2024: "Cat’s got a purrfect balance 😂". 	https://scontent-ssn1-1.cdninstagram.com/v/t51.29350-15/445364079_1148969329677809_198506909025580715_n.jpg?stp=cmp1_dst-jpg_s640x640&_nc_cat=109&ccb=1-7&_nc_sid=18de74&_nc_ohc=Pz9Pfpcq3JgQ7kNvgFTxHRY&_nc_ht=scontent-ssn1-1.cdninstagram.com&oh=00_AYBx56q-fXhStPv57TY3SLUK302gvfIU88QqEGoQlxKegg&oe=669C1F5D
2373	https://www.instagram.com/reel/C9Mp68uIJoU/?igsh=MXYzODF5b3htcTI3dg==	Lilpotates on Instagram: "Rawr 🦖 \n\n#kawaii #cute #depression #anxiety #cuteart #animated #caffeine #uwumemes #procreate #potato #procreateanimation #cuteanimation #icedcoffee #cutereels #mentalhealth #love #relationships #cutepeople #friendship #funnycomic"	15K likes, 53 comments - lilpotates on July 9, 2024: "Rawr 🦖 \n\n#kawaii #cute #depression #anxiety #cuteart #animated #caffeine #uwumemes #procreate #potato #procreateanimation #cuteanimation #icedcoffee #cutereels #mentalhealth #love #relationships #cutepeople #friendship #funnycomic". 	https://scontent-ssn1-1.cdninstagram.com/v/t51.29350-15/450358023_1982434525526056_5901834900439219531_n.jpg?stp=cmp1_dst-jpg_s640x640&_nc_cat=106&ccb=1-7&_nc_sid=18de74&_nc_ohc=lB7zxM3lKxcQ7kNvgGZvXA3&_nc_ht=scontent-ssn1-1.cdninstagram.com&oh=00_AYATwmJh2pTvOsaV74TZTY91y-kMuxaj0Uo0Jf8b-nkiGg&oe=669C1C6D
2374	https://youtube.com/shorts/YrUe-XuQRX4?si=Gri2cIkKDXRbiHs4	Japanese is so indirect that sometimes, Japanese peope don’t get it 😂: ご遠慮します	Japanese communication is sometimes too indirect or softspoken that even Japanese people don’t get it.For example, you’ll sometimes see this sign in a museum...	https://i.ytimg.com/vi/YrUe-XuQRX4/oar2.jpg?sqp=-oaymwEkCJUDENAFSFqQAgHyq4qpAxMIARUAAAAAJQAAyEI9AICiQ3gB&rs=AOn4CLAxyBUFkoFv-qIGUsAguo2sPbfmxg
2375	https://www.expatkidskorea.com/article/26-dos-and-donts-when-finding-housing-in-korea.html	26 Dos and Don’ts When Finding Housing in Korea | Expat Kids Korea: For Children and Families in Seoul Korea	26 Dos and Don’ts When Finding Housing in Korea	https://www.expatkidskorea.com/media/cache/logo_share/custom/domain_1/content_files/img_logo.png
2380	https://youtu.be/4yDm6xNeYas?si=13rUIiolDH-nvV-V	Some bad code just broke a billion Windows machines	Cybersecurity firm Crowdstrike pushed an update that caused millions of Windows computers to enter recovery mode, triggering the blue screen of death. Learn ...	https://i.ytimg.com/vi/4yDm6xNeYas/maxresdefault.jpg?sqp=-oaymwEmCIAKENAF8quKqQMa8AEB-AH-CYAC0AWKAgwIABABGGUgSChAMA8=&rs=AOn4CLBsukFcH5RItGFrSer3Pj-qky8bbQ
2381	https://httpie.io/	HTTPie – API testing client that flows with you	Making APIs simple and intuitive for those building the tools of our time.	https://httpie.io/Images/Share/default.png
2382	https://www.linode.com/docs/guides/how-to-use-fzf/	How to Install and Use fzf on Linux	Learn how to use fzf, a command-line fuzzy finder that integrates with numerous tools to improve your searches.	https://www.linode.com/docs/media/images/default_social_image.png
2383	https://github.com/sharkdp/bat?tab=readme-ov-file	GitHub - sharkdp/bat: A cat(1) clone with wings.	A cat(1) clone with wings. Contribute to sharkdp/bat development by creating an account on GitHub.	https://repository-images.githubusercontent.com/130464961/20727580-dd13-11e9-8f03-0789a00a3b64
2387	https://www.youtube.com/watch?v=mmqDYw9C30I&list=TLPQMjAwNzIwMjS-s3FLeIBCeg&index=4	7 Amazing CLI Tools You Need To Try	These are 7 game-changing cli tools for macOs or Linux operating systems. I've been incorporating them into my workflow recently and they are incredibly help...	https://i.ytimg.com/vi/mmqDYw9C30I/maxresdefault.jpg
2393	https://youtu.be/E2jrWqBDilM?si=qq02OQWV2cob6Bey	Topic Marker, Subject Marker, Object Marker (은/는, 이/가, 을/를) | Korean FAQ	It's time for a particle review of the three Korean markers - the Topic Marker (은/는), the Subject Marker (이/가), and the Object Marker (을/를). Let's compare an...	https://i.ytimg.com/vi/E2jrWqBDilM/maxresdefault.jpg
2395	https://youtu.be/7ukojv7n34Y?si=RciZm4VkroshCJip	Verb Endings ~네(요), ~지(요), ~나(요), ~군(요) | Live Class Abridged	This is an abridged version of the live stream from 3-6-2022 about the verb endings ~네(요), ~지(요), ~나(요) or ~가(요), ~(는)군(요), and ~(는)구나.Want to start learning...	https://i.ytimg.com/vi/7ukojv7n34Y/maxresdefault.jpg
2397	https://youtu.be/z_UdCUzUV4s?si=UH2iqkewyIkq9x4_	Verb Endings ~거든(요), 잖아(요), ~고(요) | Live Class Abridged	This is an abridged version of the live stream from 3-20-2022 about the verb endings ~거든(요), ~잖아(요), and ~고(요). These endings are used for a variety of situa...	https://i.ytimg.com/vi/z_UdCUzUV4s/maxresdefault.jpg
2401	https://youtube.com/shorts/6EyoRwDBSBk?si=r5ufyUK6i_pNrBiY	5 Painful Exercises That Actually Work	52 Week Guitar Player is closed for enrollment until mid-September. We will only be accepting 100 studentsCheck out the related video "How Good is 52 Week Gu...	https://i.ytimg.com/vi/6EyoRwDBSBk/oar2.jpg?sqp=-oaymwEkCJUDENAFSFqQAgHyq4qpAxMIARUAAAAAJQAAyEI9AICiQ3gB&rs=AOn4CLCGr5iV7gbrFviXBjj00z_9szRTIg
2409	https://www.youtube.com/watch?v=xp8Fzmc2upI	Valkyrae COMFORTING Miyoung	Miyoung opens up about her cat, puppy.Credits: https://www.twitch.tv/kkataminaThank you for watching :DPlease leave a like and hit the subscribe button.#miyo...	https://i.ytimg.com/vi/xp8Fzmc2upI/hqdefault.jpg
2410	https://youtube.com/shorts/X-POvzlec04?si=qv916QQ7stvHXEDq	Toast Comes Into Miyoung's Chat After Getting BANNED on Twitch #shorts #miyoung #kkatamina #toast	twitch.tv/kkatamina	https://i.ytimg.com/vi/X-POvzlec04/hq720_2.jpg?sqp=-oaymwEYCPsDENAFSFryq4qpAwoIARUAAIhC0AEB&rs=AOn4CLCRv6GtQsKqI49NjSdfsQ0gm3uWgg
2411	https://youtu.be/u482tEuWSEQ?si=2-txnzT0AfI3omT6	This game Messed Up our Friendship... Ft. Disguised Toast, Valkyrae, and Sykkuno	boku no pico park will never be the same..Peeps in the video: @DisguisedToast @LilyPichu @Valkyrae @fuslie @yvonnie @Sykkuno • Twitch - https://www.twitch.tv...	https://i.ytimg.com/vi/u482tEuWSEQ/maxresdefault.jpg
2412	https://youtu.be/dfR_LdA3fPI?si=TaLvZ0BvtS76S5U1	7 Easy Chicken Dinners	Customize & buy the Tasty Cookbook here: http://bzfd.it/2fpfeu5Check us out on Facebook! - facebook.com/buzzfeedtastyVisit these links for the recipes:One-Po...	https://i.ytimg.com/vi/dfR_LdA3fPI/maxresdefault.jpg
2414	https://youtube.com/shorts/3S2kEI18TJA?si=NYnQxqKythu1UtJA	Is she calling you boo?	#quarterjade #masayoshi #streamermoments	https://i.ytimg.com/vi/3S2kEI18TJA/oardefault.jpg?sqp=-oaymwEkCJUDENAFSFqQAgHyq4qpAxMIARUAAAAAJQAAyEI9AICiQ3gB&rs=AOn4CLBYMJh2aR3UXHQ9s2hcL2pxyKDc0w
2415	https://youtube.com/shorts/YMd0hxZeF9E?si=C-HRu42ZVFQpeQOg	Throwing shade at incompetent + rude people the Kyoto way	Kyoto people are black-belt champions at throwing shade at others in the most elegant of ways. Have you run into a service provider who is devastatingly inco...	https://i.ytimg.com/vi/YMd0hxZeF9E/oardefault.jpg?sqp=-oaymwEkCJUDENAFSFqQAgHyq4qpAxMIARUAAAAAJQAAyEI9AICiQ3gB&rs=AOn4CLBH1eRy-N2qU01qOI78vrfCnVlDdA
2416	https://youtube.com/shorts/76RZhY1Zn-o?si=s_f76zvlEQZPbvzf	And that's why I never did America's Got Talent... 🏳️‍🌈🎥🤣 | Gianmarco Soresi | Stand Up Comedy	🎤 Check Out my Tour Dates! 🎤Tour Dates: https://beacons.ai/gianmarcosoresi📱 Get a text next time I'm performing in your city 📱https://slkt.io/VbSf📩 If y...	https://i.ytimg.com/vi/76RZhY1Zn-o/oar2.jpg?sqp=-oaymwEkCJUDENAFSFqQAgHyq4qpAxMIARUAAAAAJQAAyEI9AICiQ3gB&rs=AOn4CLBTlFQGX5cFK2PrQ_CwR9s6H0AHOw
2417	https://youtu.be/4LtGy7AmI4o?si=45Mr2Ubp-Jl179st	Fuslie REALIZES she RAN AWAY from MOST THINGS in her LIFE	Fuslie searches the meaning of Avoidant Attachment issues and realizes that she has ran away from most things in her life that scared her, including her ther...	https://i.ytimg.com/vi/4LtGy7AmI4o/hqdefault.jpg
2418	https://youtu.be/lZBejB7l93c?si=scUNOJnPYKiIWTGX	Chokehold (Bluegrass Cover)	parody cover of Sleep Token's Chokehold-  ya'll worshipAll copyrights, images, and music belong to Sleep Token- no copyright infringement intended.	https://i.ytimg.com/vi/lZBejB7l93c/maxresdefault.jpg?sqp=-oaymwEmCIAKENAF8quKqQMa8AEB-AHOBYACgAqKAgwIABABGGUgWChLMA8=&rs=AOn4CLCvwSw-2barG9EeHe24aSUsfBaC8w
2419	https://www.instagram.com/reel/C742N-EOHT6/?igsh=amwxdWxtenVlbHBk	Snappy on Instagram: "It’s okay to eat alone 🙃\n\n#FoodieFinds #TorontoEats #TasteTest #TodoToronto #Canada #foodies #chickenplus #chicken #friedchicken #koreanfriedchicken"	6,205 likes, 121 comments - gosnappy.io on June 6, 2024: "It’s okay to eat alone 🙃\n\n#FoodieFinds #TorontoEats #TasteTest #TodoToronto #Canada #foodies #chickenplus #chicken #friedchicken #koreanfriedchicken". 	https://scontent-ssn1-1.cdninstagram.com/v/t51.29350-15/447915883_1132696547784728_5803116215175383107_n.jpg?stp=cmp1_dst-jpg_s640x640&_nc_cat=109&ccb=1-7&_nc_sid=18de74&_nc_ohc=1uelDJxlJrkQ7kNvgGa7zme&_nc_ht=scontent-ssn1-1.cdninstagram.com&oh=00_AYAnWjdPPwZyULFcVn5M9gxBqthFbIwvvYiD3VjrXI3K7Q&oe=66A67004
2420	https://www.instagram.com/reel/C8ClIW7KP3x/?igsh=MTN3Mzhka3V0Z255cA==	paws hub on Instagram: "Mini King 👑 this page is for sale for onek USDdee\nNo problem! Here's the information about the Mercedes CLR GTR:\n\nThe Mercedes CLR GTR is a remarkable racing car celebrated for its outstanding performance and sleek design. Powered by a potent 6.0-liter V12 engine, it delivers over 600 horsepower.\n\nAcceleration from 0 to 100 km/h takes approximately 3.7 seconds, with a remarkable top speed surpassing 320 km/h.\n\nIncorporating advanced aerodynamic features and cutting-edge stability technologies, the CLR GTR ensures exceptional stability and control, particularly during high-speed maneuvers.\n\nOriginally priced around $1.5 million, the Mercedes CLR GTR is considered one of the most exclusive and prestigious racing cars ever produced.\n\nIts limited production run of just five units adds to its rarity, making it highly sought after by racing enthusiasts and collectors worldwide. S\n\nJoin my Instagram, there are regular quizzes about cars and automotive news\n\nI invite you to my Telegram channel where you c\n\n#cutecat #cutecats #cutecatsofinstagram #cutecatofinstagram #sillycat #sillycats #sillycatsofinstagram #sillycatofinstagram #supercute #supersilly #explore #explorepage #explorepage #catsofinstagram #catofinstagram #catoftheday"	pawshub_44 on June 10, 2024: "Mini King 👑 this page is for sale for onek USDdee\nNo problem! Here's the information about the Mercedes CLR GTR:\n\nThe Mercedes CLR GTR is a remarkable racing car celebrated for its outstanding performance and sleek design. Powered by a potent 6.0-liter V12 engine, it delivers over 600 horsepower.\n\nAcceleration from 0 to 100 km/h takes approximately 3.7 seconds, with a remarkable top speed surpassing 320 km/h.\n\nIncorporating advanced aerodynamic features and cutting-edge stability technologies, the CLR GTR ensures exceptional stability and control, particularly during high-speed maneuvers.\n\nOriginally priced around $1.5 million, the Mercedes CLR GTR is considered one of the most exclusive and prestigious racing cars ever produced.\n\nIts limited production run of just five units adds to its rarity, making it highly sought after by racing enthusiasts and collectors worldwide. S\n\nJoin my Instagram, there are regular quizzes about cars and automotive news\n\nI invite you to my Telegram channel where you c\n\n#cutecat #cutecats #cutecatsofinstagram #cutecatofinstagram #sillycat #sillycats #sillycatsofinstagram #sillycatofinstagram #supercute #supersilly #explore #explorepage #explorepage #catsofinstagram #catofinstagram #catoftheday". 	https://scontent-ssn1-1.cdninstagram.com/v/t51.29350-15/448113961_1868516796923717_4703290492485411066_n.jpg?stp=cmp1_dst-jpg_s640x640&_nc_cat=1&ccb=1-7&_nc_sid=18de74&_nc_ohc=D5CYkHjC_usQ7kNvgE_8wQ1&_nc_ht=scontent-ssn1-1.cdninstagram.com&oh=00_AYCUQOlOVcNLQGhTalnSO9Blo6uyWX7cB3ayMcU_6qiucg&oe=66A67309
2421	https://youtube.com/shorts/L1dxjRAC5U8?si=6CFcVz-QfCrPMIrs	Men Are So Easily Entertained!	Watch me LIVE on TWITCH → https://twitch.tv/codemiko► 12PM PST/3PM EST/8PM GMT / ALMOST EVERYDAY!JOIN THE COMMUNITY! → https://discord.gg/codemikoTwitch Subs...	https://i.ytimg.com/vi/L1dxjRAC5U8/oardefault.jpg?sqp=-oaymwEkCJUDENAFSFqQAgHyq4qpAxMIARUAAAAAJQAAyEI9AICiQ3gB&rs=AOn4CLAy5VJ_WDgzJ4bY0UausJcdHxVOTw
2428	https://youtube.com/shorts/dy3lBL6ONIY?si=00V-lK-0WuwQ-QSR	Miyoung tries to bait Valkyrae	Subscribe here ► http://bit.ly/Sub2RaeCheck out my other socials!TikTok ► https://www.tiktok.com/@valkyrae?lang=enTwitter ► http://twitter.com/valkyraeInstag...	https://i.ytimg.com/vi/dy3lBL6ONIY/oar2.jpg?sqp=-oaymwEkCJUDENAFSFqQAgHyq4qpAxMIARUAAAAAJQAAyEI9AICiQ3gB&rs=AOn4CLBxr_9gmckVfiS7tPR_nTsOsffQRw
2436	https://www.instagram.com/p/C941nD6TRH7/?utm_source=ig_web_copy_link	CHOU_TZUYU on Instagram: "TZUYU departure to Japan 🛫 \n\n#tzuyu #choutzuyu #twice #kpop #twicetzuyu #twiceedit #trending #airport #nayeon #jihyo"	137 likes, 0 comments - tzuyusta_ on July 26, 2024: "TZUYU departure to Japan 🛫 \n\n#tzuyu #choutzuyu #twice #kpop #twicetzuyu #twiceedit #trending #airport #nayeon #jihyo". 	https://scontent-ssn1-1.cdninstagram.com/v/t51.29350-15/452945239_838320415145861_2241917903524926905_n.jpg?stp=dst-jpg_s640x640&_nc_cat=105&ccb=1-7&_nc_sid=18de74&_nc_ohc=Iie1KEjv3mgQ7kNvgHeqRgR&_nc_ht=scontent-ssn1-1.cdninstagram.com&oh=00_AYBCFmDpIWtfmY50o01UvXHc9cvyyTyHI5BYgmYWQ0XXHw&oe=66A9A56F
2437	https://www.youtube.com/shorts/l2klNbhTj7I	This Is Miyoung	thanks for watching! 🌻 social media ►►►twitter: https://twitter.com/fuslieinstagram: https://instagram.com/fusliereddit: https://www.reddit.com/r/fusliedisc...	https://i.ytimg.com/vi/l2klNbhTj7I/oar2.jpg?sqp=-oaymwEkCJUDENAFSFqQAgHyq4qpAxMIARUAAAAAJQAAyEI9AICiQ3gB&rs=AOn4CLAvgSE73eVutjtsWVUSVT2YswoINg
2440	https://www.youtube.com/watch?v=k0nqIscX2Jk	Miyoung's Comment Catches Everybody Off Guard	tina party was just differentMiyoung | https://www.twitch.tv/kkataminaLeslie | https://youtube.com/fuslieMiz | https://twitch.tv/mizkifSqueex | https://twitc...	https://i.ytimg.com/vi/k0nqIscX2Jk/maxresdefault.jpg
2442	https://gimmedelicious.com/quick-easy-15-minute-chicken-pasta/	15 minute Chicken Pasta	Chicken and pasta cooked in a garlic tomato sauce and topped with mozzarella & parmesan cheese cooked in just 15 minutes and will become your new favorite weeknight meal! Today I'm going to share a	https://gimmedelicious.com/wp-content/uploads/2019/12/15-Minute-Chicken-Pasta-3.jpg
2443	https://youtube.com/shorts/0uSfZKWoFjM?si=gTrvWa4Fh_2A3k2z	How Many Vocal Layers Are in This Song? #iprevail #shorts #music		https://i.ytimg.com/vi/0uSfZKWoFjM/hq720.jpg?sqp=-oaymwEkCJUDENAFSFryq4qpAxYIARUAAAAAJQAAyEI9AICiQ3gB0AEB&rs=AOn4CLDRHDGKUX3v3cTxJml_KEbfW48nhQ
2449	https://youtube.com/shorts/YvFdgDgnn68?si=0vY8UK3iYhjjTjMR	Baking Cookies During a Concert		https://i.ytimg.com/vi/YvFdgDgnn68/oardefault.jpg?sqp=-oaymwEkCJUDENAFSFqQAgHyq4qpAxMIARUAAAAAJQAAyEI9AICiQ3gB&rs=AOn4CLDLpF7RJyv_OJCcBNbmhBqbUoKAlg
2450	https://ko.aliexpress.com/item/1005002709729784.html?spm=a2g0o.productlist.main.5.1f5byhwmyhwmbd&algo_pvid=73249568-ef57-4c36-b09e-9ff72083c22d&algo_exp_id=73249568-ef57-4c36-b09e-9ff72083c22d-2&pdp_npi=4%40dis%21KRW%2129047%2113078%21%21%2120.39%219.18%21%402141113f17221495258587655ed261%2112000021806380768%21sea%21KR%210%21ABX&curPageLogUid=LaRKZ4TyMlFB&utparam-url=scene%3Asearch%7Cquery_from%3A	중국 매화 스크롤 걸이식 캔버스 포스터, 빈티지 꽃 벽 그림 스크롤 페인팅, 거실 사무실 장식 - AliExpress 15	Smarter Shopping, Better Living! Aliexpress.com	https://ae01.alicdn.com/kf/H5af45826a2204a18bad8ec404a540f055/-.jpg
2451	https://ko.aliexpress.com/item/1005006948298260.html?srcSns=sns_Copy&spreadType=socialShare&bizType=ProductDetail&social_params=6000127768541&aff_fcid=c472b2958f9a4131ab1915b490a3ba96-1722149359948-07239-_oDNZB60&tt=MG&aff_fsk=_oDNZB60&aff_platform=default&sk=_oDNZB60&aff_trace_key=c472b2958f9a4131ab1915b490a3ba96-1722149359948-07239-_oDNZB60&shareId=6000127768541&businessType=ProductDetail&platform=AE&terminal_id=8b4277741eb5472f8835eec3ff1150b8&afSmartRedirect=y&gatewayAdapt=glo2kor	남성용 쿨 화이트 블랙 디자인 패션, 스카이 림 인스파이어드, 전설적인 게임 패턴, 탑 티셔츠 - AliExpress 200000345	Smarter Shopping, Better Living! Aliexpress.com	https://ae01.alicdn.com/kf/S5663db565b2e4a308b80d8da404f88461/-.jpg
2454	https://youtube.com/shorts/7seSIKMZbPM?si=rULjF9BSxViEBdWS	Knowledge Transfer	Recorded live on twitch, GET IN https://twitch.tv/ThePrimeagenBecome a backend engineer.  Its my favorite sitehttps://boot.dev?promo=PRIMEYTThis is also the ...	https://i.ytimg.com/vi/7seSIKMZbPM/oar2.jpg?sqp=-oaymwEkCJUDENAFSFqQAgHyq4qpAxMIARUAAAAAJQAAyEI9AICiQ3gB&rs=AOn4CLCYqzzq2oc8qcmV1UJSYJJMdvBOUg
2455	https://youtube.com/shorts/lL8lA4BecHs?si=BFYyrszuCdhCdRsi	Dungeons and Dragons - Computer Edition	DnD was the original Client/Server game. ( ͡° ͜ʖ ͡°)Watch the stream here:https://piratesoftware.live#Shorts #Twitch #PirateSoftware	https://i.ytimg.com/vi/lL8lA4BecHs/oar2.jpg?sqp=-oaymwEkCJUDENAFSFqQAgHyq4qpAxMIARUAAAAAJQAAyEI9AICiQ3gB&rs=AOn4CLAAJ6seKgZpSwTNGd_NH2YJuv0L_w
2460	https://youtu.be/BdOtR2A5agc?si=mUSQr0qwXDBfy06e	How rap sounds to non-english speakers	An unusual word takes over the spelling bee.IGOWALLAH (ft. Hoodie Guy)directed by Daniel Thrasherfeaturing   @williamosman  music by oo oxygencinematography ...	https://i.ytimg.com/vi/BdOtR2A5agc/maxresdefault.jpg
2463	https://youtu.be/B7eKYJwsxqc?si=GyRGp1j48dh0sfuf	getting wet with @ExtraEmily	fanfan twitch stream clips and highlights — watch me live at https://www.twitch.tv/fanfan :)open me!!follow me:➤ twitch - https://www.twitch.tv/fanfan➤ twitt...	https://i.ytimg.com/vi/B7eKYJwsxqc/maxresdefault.jpg
2468	https://youtu.be/49H7054w10I?si=dgEM1pHicLJNoNMH	Gen Z Courtroom	WATCH BLOOPERS FOR EVERY VIDEO ▶ https://www.patreon.com/stevieemersonGET SOME MERCH ▶ https://slappersonly.shop/FOLLOW ME:INSTAGRAM ▶ https://www.instagram....	https://i.ytimg.com/vi/49H7054w10I/hqdefault.jpg
2470	https://youtube.com/shorts/0u2jN5TJINI?si=-x9-31vDL328ycRf	FUSLIE PROTECTS MIYOUNG FROM CHAT	thanks for watching! 🌻 social media ►►►twitter: https://twitter.com/fuslieinstagram: https://instagram.com/fusliereddit: https://www.reddit.com/r/fusliedisc...	https://i.ytimg.com/vi/0u2jN5TJINI/oar2.jpg?sqp=-oaymwEkCJUDENAFSFqQAgHyq4qpAxMIARUAAAAAJQAAyEI9AICiQ3gB&rs=AOn4CLB5hOwTIF2XbKsb8SweVHSMp2GYHQ
2471	https://youtube.com/shorts/EaezX2qGAak?si=YGmDWY__5DK0t_Yd	Tina drank a lil too much	Subscribe here ► http://bit.ly/Sub2RaeCheck out my other socials!TikTok ► https://www.tiktok.com/@valkyrae?lang=enTwitter ► http://twitter.com/valkyraeInstag...	https://i.ytimg.com/vi/EaezX2qGAak/oar2.jpg?sqp=-oaymwEkCJUDENAFSFqQAgHyq4qpAxMIARUAAAAAJQAAyEI9AICiQ3gB&rs=AOn4CLDlJ9AWTNEQJjTODgonJc1_FmWyHg
2472	https://youtu.be/PsdupJ2qy1M?si=xVQusABLtKQHNMBg	꽁꽁 얼어붙은 수박위로 🍉 츄대표가 걸어다닙니다 💨[롯대행사] ep.1 여름 이색 메뉴 개발기 (별점이 생명★★★★★)	✉의뢰인 : 카페 업고 튀어 [롯대행사] ep.1 ✉🗂롯대행사의 첫 심부름 의뢰🗂📢츄대표님 ~! 저희 가게 대박 가게로 거듭날 수 있게 도와주세요!👀Sure. Why not~? 오 ~ 롯데마트 ~ 롯데마트 ~츄대표와 괴식요리사 풍인턴이 만든 특별한 여름 이색 메뉴꽁꽁 얼어붙...	https://i.ytimg.com/vi/PsdupJ2qy1M/maxresdefault.jpg
2474	https://youtu.be/UZfi4VueGM4?si=0__kqogdNikJo7zi	Tina Being Awkwardly Adorable for 5 Minutes	One of the most UNDERRATED FEMALE STREAMERS on TWITCH!TinaKitten (YouTube): https://www.youtube.com/channel/UC6gED92IrqY4S94rCC0Z3SQTinaKitten (Twitch): http...	https://i.ytimg.com/vi/UZfi4VueGM4/maxresdefault.jpg
2476	https://youtube.com/shorts/2EdcjBBA4MM?si=UMCuLNbzkG-HNBN_	Tinakitten misheard Valkyrae	Check out my other socials!Valkyrae ► https://www.youtube.com/c/Valkyrae1Valkyrae2 ► https://www.youtube.com/c/Valkyrae2Twitter ► http://twitter.com/valkyrae...	https://i.ytimg.com/vi/2EdcjBBA4MM/oar2.jpg?sqp=-oaymwEkCJUDENAFSFqQAgHyq4qpAxMIARUAAAAAJQAAyEI9AICiQ3gB&rs=AOn4CLD1k8u8qrGQa6GXm91LzQwch2rTIA
2487	https://youtube.com/shorts/1ilUpXkGxbo?si=F9GU4dSbzS1XUjG0	Email Mike	Watch the stream here:https://piratesoftware.liveJoin the community here:https://discord.gg/piratesoftware#Shorts #Twitch #PirateSoftware	https://i.ytimg.com/vi/1ilUpXkGxbo/oar2.jpg?sqp=-oaymwEkCJUDENAFSFqQAgHyq4qpAxMIARUAAAAAJQAAyEI9AICiQ3gB&rs=AOn4CLB4Tut-aOoIy-NDj1m5dQQqY6bILQ
2489	https://youtube.com/shorts/RyQcPAOQbV8?si=Qm-96JlrSW8XtS9p	does size matter?	thanks for watching! ✨*my socials*FAVES - https://amzn.to/3N2hO9rTWITCH - http://twitch.tv/xchocobarsTWITTER - https://twitter.com/xchocobarsINSTAGRAM - http...	https://i.ytimg.com/vi/RyQcPAOQbV8/oardefault.jpg?sqp=-oaymwEkCJUDENAFSFqQAgHyq4qpAxMIARUAAAAAJQAAyEI9AICiQ3gB&rs=AOn4CLDDeK7f4TkdcKXR8z9O3nYgTgxAUA
2491	https://youtube.com/shorts/3amaXasj2-I?si=hBhv8A-LM5c1jcnL	he thought he had rizz	thanks for watching! ✨*my socials*FAVES - https://amzn.to/3N2hO9rTWITCH - http://twitch.tv/xchocobarsTWITTER - https://twitter.com/xchocobarsINSTAGRAM - http...	https://i.ytimg.com/vi/3amaXasj2-I/oardefault.jpg?sqp=-oaymwEkCJUDENAFSFqQAgHyq4qpAxMIARUAAAAAJQAAyEI9AICiQ3gB&rs=AOn4CLArZDsehZAL6-XXOFmZa5Rm2TgFjw
2492	https://youtube.com/shorts/LnGbaNq58es?si=TjsHQrKzfYPtzXRG	i didn't realize everyone could hear me...	thanks for watching! ✨*my socials*FAVES - https://amzn.to/3N2hO9rTWITCH - http://twitch.tv/xchocobarsTWITTER - https://twitter.com/xchocobarsINSTAGRAM - http...	https://i.ytimg.com/vi/LnGbaNq58es/oardefault.jpg?sqp=-oaymwEkCJUDENAFSFqQAgHyq4qpAxMIARUAAAAAJQAAyEI9AICiQ3gB&rs=AOn4CLDr3Yk9GXI4t3o4M6yz1a4gtronSQ
2493	https://youtube.com/shorts/ITPOSriLjtM?si=r6jz-XK7VDuPa5eP	How to keep your stir-fry chicken tender and juicy 😋	#shorts #cooking #food #chinese #recipe #chicken #stirfry #chef	https://i.ytimg.com/vi/ITPOSriLjtM/oar2.jpg?sqp=-oaymwEkCJUDENAFSFqQAgHyq4qpAxMIARUAAAAAJQAAyEI9AICiQ3gB&rs=AOn4CLB-Lp3fLnKnaGSWF1APZNmjf9xmig
2495	https://youtube.com/shorts/Z59_raChnPA?si=nt4baXPl5ImzEl5E	I Make Square Eggs ⬜️	This video shows you how to turn a boiled egg into a square or cubed boiled egg, some may call them Minecraft eggs.Using a special gadget that turns boiled e...	https://i.ytimg.com/vi/Z59_raChnPA/oardefault.jpg?sqp=-oaymwEkCJUDENAFSFqQAgHyq4qpAxMIARUAAAAAJQAAyEI9AICiQ3gB&rs=AOn4CLBQHFF-_jA2K-7RXXh2O8xCoZyNZw
2500	https://youtube.com/shorts/Jx8VoKmpW5w?si=wLK3WmbNjzqtAjkN	finding out news in 2024		https://i.ytimg.com/vi/Jx8VoKmpW5w/oardefault.jpg?sqp=-oaymwEkCJUDENAFSFqQAgHyq4qpAxMIARUAAAAAJQAAyEI9AICiQ3gB&rs=AOn4CLCbhVQUwaSbLrwpAm3LP8N0jw5EGQ
2506	https://youtu.be/wqAvHz5lwZ8?si=wV8bgzc2u0yEnmAH	JLPT N4  180 Kanji animation Step 4 -  No. 71 - 90	Those who have been waiting for this video, thank you for waiting! Finally I could finish making and upload it! Phew!Especially those who are having JLPT N4 ...	https://i.ytimg.com/vi/wqAvHz5lwZ8/hqdefault.jpg
2507	https://youtu.be/GQVVQq_wk4Y?si=n9Caq1gnE4GcleS8	Janet learns how FREAKY Fanfan is	VOD: https://youtu.be/h9aLOOFfalU?si=YzUTiULZr3u4VoxnPlease subscribe: https://goo.gl/3KNHZF	https://i.ytimg.com/vi/GQVVQq_wk4Y/maxresdefault.jpg
2555	https://youtu.be/OiPPlQ1yWr4?si=Hka7YpgTnYRL9w3U	A Small Update On The SITUATION	▼▼ Expand for MORE SHIA SURPRISES! ▼▼► Twitch: https://www.twitch.tv/potasticp► Twitter: https://twitter.com/potasticpanda22► Reddit: https://www.reddit.com/...	https://i.ytimg.com/vi/OiPPlQ1yWr4/maxresdefault.jpg
2556	https://www.instagram.com/reel/C-WrfnUSPxE/?igsh=bmRkNGNyOWhiNmJm	Jin on Instagram: "HOW I MADE PEM PEM by @gmengzstillyoung ft @nonono_rith #producer #music #cambodia #phnompenh #omens"	148 likes, 12 comments - jin_avr on August 6, 2024: "HOW I MADE PEM PEM by @gmengzstillyoung ft @nonono_rith #producer #music #cambodia #phnompenh #omens". 	https://scontent-ssn1-1.cdninstagram.com/v/t51.29350-15/454454297_8140125312677074_7046066015249744049_n.jpg?stp=cmp1_dst-jpg_s640x640&_nc_cat=110&ccb=1-7&_nc_sid=18de74&_nc_ohc=ANuv4nm7SLUQ7kNvgGK_cX0&_nc_ht=scontent-ssn1-1.cdninstagram.com&oh=00_AYB5wpgOPJAc30TGyJLq8cLc2Zy1tn6wqDAJtjyRJ7pbaA&oe=66BE51CC
2558	https://youtube.com/shorts/1hcPktp_P7I?si=b5ML_H6ecDOXAZOK	this kid bro  #kkatamina #miyoungandfriends #shorts	Get 1 month of Xbox Game Pass Ultimate for every Predator product! Click the link for more info: https://acer.link/kkatamina #adThank you for watching! *:･ﾟ✧...	https://i.ytimg.com/vi/1hcPktp_P7I/oar2.jpg?sqp=-oaymwEkCJUDENAFSFqQAgHyq4qpAxMIARUAAAAAJQAAyEI9AICiQ3gB&rs=AOn4CLD4slv6JSBGdS0EMEj59QOaavC48w
2559	https://youtu.be/g4Pb25BLkKQ?si=9FyHBMQtq_I32GUV	ファタール / GEMN - 弾き語り風	むずすぎる勘弁してくださいファタール / GEMN - Fatal / GEMNGEMN(中島健人・キタニタツヤ)  “ファタール”(Fatal)TVアニメ『【推しの子】』第2期オープニング主題歌Arrange - 少年T素敵な動画お借りしました↓https://www.youtube.com/watch?v=e...	https://i.ytimg.com/vi/g4Pb25BLkKQ/maxresdefault.jpg
2561	https://youtu.be/pnWMGLGfDys?si=_SxwpqwJ4EChAv-5	luv - Fuwa Fuwa (Official Music Video)	https://luv.lnk.to/FuwaFuwaDigital Single「Fuwa Fuwa」配信中---------------------------------------------------------------------------------10月20日に初の自主企画イベント開催決定...	https://i.ytimg.com/vi/pnWMGLGfDys/maxresdefault.jpg
\.


--
-- TOC entry 3342 (class 0 OID 180251)
-- Dependencies: 217
-- Data for Name: note_tag; Type: TABLE DATA; Schema: public; Owner: flashcards_owner
--

COPY public.note_tag (note, tag) FROM stdin;
2291	algorithm
2291	concept
2317	korean
2316	korean
2316	shorts
2316	youtube
2346	korean
2346	reading
2345	korean
2345	reading
2347	korean
2347	reading
2347	decor
2387	terminal
2387	linux
2383	terminal
2383	linux
2382	terminal
2382	linux
2381	terminal
2381	linux
2393	korean
2395	korean
2397	korean
2397	guitar
2412	cooking
2421	asd
2421	meme
2420	cat
2440	meme
2440	otv
2436	kpop
2436	tzuyu
2443	metal
2442	cooking
2437	meme
2437	otv
2418	music
2419	haha
2450	shopping
2451	shopping
2428	haha
2463	haha
2460	haha
2474	haha
2476	haha
2476	shopping
2489	haha
2507	haha
2507	otv
2558	otv
2558	haha
2559	music
2506	japanese
2506	kanji
2561	music
2561	jpop
2559	jrock
\.


--
-- TOC entry 3340 (class 0 OID 172032)
-- Dependencies: 215
-- Data for Name: tag; Type: TABLE DATA; Schema: public; Owner: flashcards_owner
--

COPY public.tag (tag_name) FROM stdin;
algorithm
concept
korean
shorts
youtube
reading
decor
terminal
linux
guitar
cooking
asd
meme
cat
shopping
asdasd
otv
miyoung
kpop
tzuyu
metal
music
haha
japanese
kanji
jpop
jrock
test
new
\.


--
-- TOC entry 3349 (class 0 OID 0)
-- Dependencies: 218
-- Name: note_in_seq; Type: SEQUENCE SET; Schema: public; Owner: flashcards_owner
--

SELECT pg_catalog.setval('public.note_in_seq', 2561, true);


--
-- TOC entry 3192 (class 2606 OID 180237)
-- Name: note note_pkey; Type: CONSTRAINT; Schema: public; Owner: flashcards_owner
--

ALTER TABLE ONLY public.note
    ADD CONSTRAINT note_pkey PRIMARY KEY (id);


--
-- TOC entry 3194 (class 2606 OID 188417)
-- Name: note_tag note_tag_pk; Type: CONSTRAINT; Schema: public; Owner: flashcards_owner
--

ALTER TABLE ONLY public.note_tag
    ADD CONSTRAINT note_tag_pk PRIMARY KEY (note, tag);


--
-- TOC entry 3190 (class 2606 OID 172036)
-- Name: tag tag_pkey; Type: CONSTRAINT; Schema: public; Owner: flashcards_owner
--

ALTER TABLE ONLY public.tag
    ADD CONSTRAINT tag_pkey PRIMARY KEY (tag_name);


--
-- TOC entry 3195 (class 2606 OID 229381)
-- Name: note_tag note_tag_note_fkey; Type: FK CONSTRAINT; Schema: public; Owner: flashcards_owner
--

ALTER TABLE ONLY public.note_tag
    ADD CONSTRAINT note_tag_note_fkey FOREIGN KEY (note) REFERENCES public.note(id) ON DELETE CASCADE;


--
-- TOC entry 3196 (class 2606 OID 229376)
-- Name: note_tag note_tag_tag_fkey; Type: FK CONSTRAINT; Schema: public; Owner: flashcards_owner
--

ALTER TABLE ONLY public.note_tag
    ADD CONSTRAINT note_tag_tag_fkey FOREIGN KEY (tag) REFERENCES public.tag(tag_name) ON DELETE CASCADE;


--
-- TOC entry 2047 (class 826 OID 163842)
-- Name: DEFAULT PRIVILEGES FOR SEQUENCES; Type: DEFAULT ACL; Schema: public; Owner: cloud_admin
--

ALTER DEFAULT PRIVILEGES FOR ROLE cloud_admin IN SCHEMA public GRANT ALL ON SEQUENCES TO neon_superuser WITH GRANT OPTION;


--
-- TOC entry 2046 (class 826 OID 163841)
-- Name: DEFAULT PRIVILEGES FOR TABLES; Type: DEFAULT ACL; Schema: public; Owner: cloud_admin
--

ALTER DEFAULT PRIVILEGES FOR ROLE cloud_admin IN SCHEMA public GRANT ALL ON TABLES TO neon_superuser WITH GRANT OPTION;


-- Completed on 2024-08-12 05:44:44 UTC

--
-- PostgreSQL database dump complete
--
