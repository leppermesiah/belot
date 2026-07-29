// ---------------------------------------------------------------
// i18n - Bulgarian/English. Every user-facing string lives here so
// switching languages (lobby stacked buttons, or the in-game flag
// buttons) re-renders everything consistently. Values are either a
// plain string or a function(...args) for templated strings.
// ---------------------------------------------------------------
const I18N = {
  bg: {
    title: 'Белот',
    lobbyTitle: '♣ ♦ ♥ ♠ Белот',
    yourName: 'Твоето име',
    createGame: 'Създай игра',
    or: 'или',
    roomCodePlaceholder: 'Код на стаята',
    join: 'Присъедини се',
    waitingForPlayers: 'Чакаме играчи…',
    roomCodeLabel: 'Код на стаята:',
    chooseTeam: 'Избери отбор:',
    teamADefault: 'Отбор А',
    teamBDefault: 'Отбор Б',
    start: 'Старт',
    chooseTeamNameTitle: 'Избери име на своя отбор',
    chooseTeamNameBody: 'Ти беше избран/а да кръстиш отбора — това име ще остане до края на играта.',
    teamNamePlaceholder: 'Име на отбора',
    done: 'Готово',
    us: 'Ние',
    them: 'Вие',
    newGame: 'Нова игра',
    suitClubs: '♣ Спатия',
    suitDiamonds: '♦ Каро',
    suitHearts: '♥ Купа',
    suitSpades: '♠ Пика',
    noTrump: 'Без коз',
    allTrump: 'Всичко коз',
    pass: 'Пас',
    contra: 'Контра',
    reconto: 'Реконтра',
    announceHave: 'Имаш за обявяване:',
    declareBelotTitle: 'Обяви белот?',
    declareBelotBody: 'Държиш поп и дама коз — обявяваш ли белот преди да играеш картата?',
    yesDeclareBelot: 'Да, обявявам белот',
    no: 'Не',
    continueBtn: 'Продължи',
    reconnecting: 'Възстановяване на връзката…',
    seatLabel: (i) => `Място ${i}`,
    seatEmpty: (i) => `Място ${i}: —`,
    namedAs: (name) => `Кръстен: ${name}`,
    choosingName: 'Избира се име…',
    contraLine: (name) => `Контра (${name})`,
    recontoLine: (name) => `Реконтра (${name})`,
    withContra: ' с КОНТРА',
    withReconto: ' с РЕКОНТРА',
    youLabel: (name) => `Ти (${name})`,
    dealerSuffixParen: ' (раздавач)',
    dealerSuffixDash: ' — раздавач',
    bidPromptHasContract: 'Пасирай, обяви нещо по-силно, или контрирай/реконтрирай.',
    bidPromptNoContract: 'Твой ред е да обявиш.',
    waitingPlayersCount: (r, total) => `Чакаме останалите играчи… (${r}/${total})`,
    waitingPlayers: 'Чакаме останалите играчи…',
    gameOver: 'Играта свърши!',
    teamWinsMatch: (team) => `${team} печели мача!`,
    enterRoomCode: 'Въведи код на стаята.',
    enterYourName: 'Въведи си името.',
    confirmLeaveRoom: 'Да напуснеш тази игра и да започнеш нова стая?',
    connectingPrevious: 'Свързване към предишната игра…',
    announcePrefix: (label) => `Обяви: ${label}`,
    suitNames: { clubs: 'спатия', diamonds: 'каро', hearts: 'купа', spades: 'пика' },
    suitNamesCap: { clubs: 'Спатия', diamonds: 'Каро', hearts: 'Купа', spades: 'Пика' },
    cardStrengthTitle: 'Сила на картите',
    cardStrengthOther: 'Останали',
    rankDisplay: { A: 'Асо', K: 'Поп', Q: 'Дама', J: 'Вале' },
    annTierce: (rank, suit) => `Терца до ${rank} ${suit}`,
    annFifty: (rank, suit) => `Петдесет до ${rank} ${suit}`,
    annHundred: (rank, suit) => `Сто до ${rank} ${suit}`,
    annCarreJ: 'Каре валета',
    annCarreNine: 'Каре девятки',
    annCarreOther: (rank) => `Каре ${rank}`,
    annBelot: 'Белот',
    errorRoomNotFound: 'Стаята не е намерена.',
    errorRoomFull: 'Стаята вече е пълна.',
    errorTeamsIncomplete: 'Всеки играч трябва да избере отбор (2 Черешка, 2 Малинка) преди старт.',
    errorInvalidTeam: 'Невалиден отбор.',
    errorTeamFull: 'Този отбор вече е пълен.',
  },
  en: {
    title: 'Belot',
    lobbyTitle: '♣ ♦ ♥ ♠ Belot',
    yourName: 'Your name',
    createGame: 'Create game',
    or: 'or',
    roomCodePlaceholder: 'Room code',
    join: 'Join',
    waitingForPlayers: 'Waiting for players…',
    roomCodeLabel: 'Room code:',
    chooseTeam: 'Choose a team:',
    teamADefault: 'Team A',
    teamBDefault: 'Team B',
    start: 'Start',
    chooseTeamNameTitle: 'Choose a name for your team',
    chooseTeamNameBody: 'You were picked to name the team — this name sticks for the rest of the match.',
    teamNamePlaceholder: 'Team name',
    done: 'Done',
    us: 'Us',
    them: 'Them',
    newGame: 'New game',
    suitClubs: '♣ Clubs',
    suitDiamonds: '♦ Diamonds',
    suitHearts: '♥ Hearts',
    suitSpades: '♠ Spades',
    noTrump: 'No trump',
    allTrump: 'All trump',
    pass: 'Pass',
    contra: 'Double',
    reconto: 'Redouble',
    announceHave: 'You can announce:',
    declareBelotTitle: 'Declare belot?',
    declareBelotBody: 'You hold the King and Queen of trump — declare belot before playing the card?',
    yesDeclareBelot: 'Yes, declare belot',
    no: 'No',
    continueBtn: 'Continue',
    reconnecting: 'Reconnecting…',
    seatLabel: (i) => `Seat ${i}`,
    seatEmpty: (i) => `Seat ${i}: —`,
    namedAs: (name) => `Named: ${name}`,
    choosingName: 'Choosing a name…',
    contraLine: (name) => `Double (${name})`,
    recontoLine: (name) => `Redouble (${name})`,
    withContra: ' WITH DOUBLE',
    withReconto: ' WITH REDOUBLE',
    youLabel: (name) => `You (${name})`,
    dealerSuffixParen: ' (dealer)',
    dealerSuffixDash: ' — dealer',
    bidPromptHasContract: 'Pass, bid something higher, or double/redouble.',
    bidPromptNoContract: 'Your turn to bid.',
    waitingPlayersCount: (r, total) => `Waiting for the other players… (${r}/${total})`,
    waitingPlayers: 'Waiting for the other players…',
    gameOver: 'Game over!',
    teamWinsMatch: (team) => `${team} wins the match!`,
    enterRoomCode: 'Enter the room code.',
    enterYourName: 'Enter your name.',
    confirmLeaveRoom: 'Leave this game and start a new room?',
    connectingPrevious: 'Connecting to your previous game…',
    announcePrefix: (label) => `Announce: ${label}`,
    suitNames: { clubs: 'clubs', diamonds: 'diamonds', hearts: 'hearts', spades: 'spades' },
    suitNamesCap: { clubs: 'Clubs', diamonds: 'Diamonds', hearts: 'Hearts', spades: 'Spades' },
    cardStrengthTitle: 'Card strength',
    cardStrengthOther: 'Other suits',
    rankDisplay: { A: 'Ace', K: 'King', Q: 'Queen', J: 'Jack' },
    annTierce: (rank, suit) => `Tierce to ${rank} of ${suit}`,
    annFifty: (rank, suit) => `Fifty to ${rank} of ${suit}`,
    annHundred: (rank, suit) => `Hundred to ${rank} of ${suit}`,
    annCarreJ: 'Four jacks',
    annCarreNine: 'Four nines',
    annCarreOther: (rank) => `Four ${rank}s`,
    annBelot: 'Belot',
    errorRoomNotFound: 'Room not found.',
    errorRoomFull: 'This room is already full.',
    errorTeamsIncomplete: 'Every player must choose a team (2 Cherry, 2 Raspberry) before starting.',
    errorInvalidTeam: 'Invalid team.',
    errorTeamFull: 'That team is already full.',
  },
};

let currentLang = 'bg';
try {
  const savedLang = localStorage.getItem('belotLang');
  if (savedLang === 'bg' || savedLang === 'en') currentLang = savedLang;
} catch (e) { /* storage unavailable - default to bg */ }

function t(key, ...args) {
  const dict = I18N[currentLang] || I18N.bg;
  const val = dict[key];
  if (typeof val === 'function') return val(...args);
  return val !== undefined ? val : key;
}

let lastRoomState = null;

function setLang(lang) {
  if (lang !== 'bg' && lang !== 'en') return;
  currentLang = lang;
  try { localStorage.setItem('belotLang', lang); } catch (e) { /* not fatal */ }
  document.documentElement.lang = lang;
  applyStaticTranslations();
  if (lastState) renderGameState(lastState);
  if (lastRoomState) onRoomState(lastRoomState);
}

function applyStaticTranslations() {
  document.getElementById('pageTitle').textContent = t('title');
  document.getElementById('lobbyTitle').textContent = t('lobbyTitle');
  document.getElementById('nameInput').placeholder = t('yourName');
  document.getElementById('createBtn').textContent = t('createGame');
  document.getElementById('orDivider').textContent = t('or');
  document.getElementById('codeInput').placeholder = t('roomCodePlaceholder');
  document.getElementById('joinBtn').textContent = t('join');
  document.getElementById('waitingTitle').textContent = t('waitingForPlayers');
  document.getElementById('chooseTeamLabel').textContent = t('chooseTeam');
  document.getElementById('startBtn').textContent = t('start');
  document.getElementById('teamNameTitle').textContent = t('chooseTeamNameTitle');
  document.getElementById('teamNameBody').textContent = t('chooseTeamNameBody');
  document.getElementById('teamNameInput').placeholder = t('teamNamePlaceholder');
  document.getElementById('teamNameDoneBtn').textContent = t('done');
  document.getElementById('leaveRoomBtn').textContent = t('newGame');
  document.getElementById('bidClubsBtn').textContent = t('suitClubs');
  document.getElementById('bidDiamondsBtn').textContent = t('suitDiamonds');
  document.getElementById('bidHeartsBtn').textContent = t('suitHearts');
  document.getElementById('bidSpadesBtn').textContent = t('suitSpades');
  document.getElementById('noTrumpBtn').textContent = t('noTrump');
  document.getElementById('allTrumpBtn').textContent = t('allTrump');
  document.getElementById('passBtn').textContent = t('pass');
  document.getElementById('contraBtn').textContent = t('contra');
  document.getElementById('recontoBtn').textContent = t('reconto');
  document.getElementById('announceHaveLabel').textContent = t('announceHave');
  document.getElementById('belotTitle').textContent = t('declareBelotTitle');
  document.getElementById('belotBody').textContent = t('declareBelotBody');
  document.getElementById('belotYesBtn').textContent = t('yesDeclareBelot');
  document.getElementById('belotNoBtn').textContent = t('no');
  document.getElementById('recapDoneBtn').textContent = t('done');
  document.getElementById('resultCloseBtn').textContent = t('continueBtn');
  document.getElementById('newGameBtn').textContent = t('newGame');
  document.getElementById('connStatus').textContent = t('reconnecting');
  document.getElementById('chooseCherryBtn').textContent = t('teamADefault');
  document.getElementById('chooseMalinaBtn').textContent = t('teamBDefault');
  if (currentRoomCode) renderRoomCodeLabel();

  const bgActive = currentLang === 'bg';
  document.getElementById('langBgBtn')?.classList.toggle('active-lang', bgActive);
  document.getElementById('langEnBtn')?.classList.toggle('active-lang', !bgActive);
  document.getElementById('langFlagBgBtn')?.classList.toggle('active-lang', bgActive);
  document.getElementById('langFlagEnBtn')?.classList.toggle('active-lang', !bgActive);
}

document.getElementById('langBgBtn')?.addEventListener('click', () => setLang('bg'));
document.getElementById('langEnBtn')?.addEventListener('click', () => setLang('en'));
document.getElementById('langFlagBgBtn')?.addEventListener('click', () => setLang('bg'));
document.getElementById('langFlagEnBtn')?.addEventListener('click', () => setLang('en'));

const SUIT_SYMBOL = { clubs: '♣', diamonds: '♦', hearts: '♥', spades: '♠' };
const SUIT_RED = { diamonds: true, hearts: true };
const SUIT_LETTER = { clubs: 'C', diamonds: 'D', hearts: 'H', spades: 'S' };

// Left-to-right suit order for sorting your own hand.
const HAND_SUIT_ORDER = ['spades', 'hearts', 'diamonds', 'clubs'];
// Rank order within a suit depends on whether that suit counts as trump
// under the current contract (всичко коз = every suit; без коз = none;
// suit contract = only the called suit). Matches engine/card.go's
// nonTrumpRankOrder/trumpRankOrder (low->high) - this is real trick-taking
// strength, not an arbitrary display order.
const NORMAL_RANK_ORDER = ['7', '8', '9', 'J', 'Q', 'K', '10', 'A'];
const TRUMP_RANK_ORDER = ['7', '8', 'Q', 'K', '10', 'A', '9', 'J'];

// Strongest-to-weakest, for the "card strength" hint.
const NORMAL_STRENGTH_ORDER = [...NORMAL_RANK_ORDER].reverse();
const TRUMP_STRENGTH_ORDER = [...TRUMP_RANK_ORDER].reverse();

function isTrumpSuit(suit, m) {
  if (!m.contract) return false;
  if (m.contract.type === 'alltrump') return true;
  if (m.contract.type === 'notrump') return false;
  return suit === m.trump;
}

function sortHand(hand, m) {
  // The suit that's actually trump right now uses the trump strength
  // order (J,9,A,10,K,Q,8,7); every other suit uses the normal order
  // (A,10,K,Q,J,9,8,7). Under всичко коз every suit counts as trump;
  // under без коз or before any contract is set, none does. Sorted
  // strongest to weakest, left to right.
  return [...(hand || [])].sort((a, b) => {
    const ca = parseCard(a), cb = parseCard(b);
    const suitDiff = HAND_SUIT_ORDER.indexOf(ca.suit) - HAND_SUIT_ORDER.indexOf(cb.suit);
    if (suitDiff !== 0) return suitDiff;
    const order = isTrumpSuit(ca.suit, m) ? TRUMP_RANK_ORDER : NORMAL_RANK_ORDER;
    return order.indexOf(cb.rank) - order.indexOf(ca.rank);
  });
}

// Returns the trump suit for which this hand holds BOTH King and Queen
// (belot-eligible), or null. Never applies under без коз (no trump).
// Under всичко коз specifically, every suit counts as trump, but belot
// only actually counts if you're leading the trick (nothing played
// yet) or the suit matches whatever's currently led - playing an
// off-suit K/Q there (forced discard because you're out of the led
// suit) does NOT count, even though that suit is technically trump.
function findBelotSuit(hand, m) {
  const bySuit = {};
  (hand || []).forEach(cardStr => {
    const { rank, suit } = parseCard(cardStr);
    if (!bySuit[suit]) bySuit[suit] = {};
    bySuit[suit][rank] = true;
  });
  const isAllTrump = m.contract && m.contract.type === 'alltrump';
  const ledSuit = (m.currentTrick && m.currentTrick.length > 0) ? parseCard(m.currentTrick[0].card).suit : null;
  for (const suit of Object.keys(bySuit)) {
    if (!bySuit[suit]['K'] || !bySuit[suit]['Q']) continue;
    if (!isTrumpSuit(suit, m)) continue;
    if (isAllTrump && ledSuit !== null && suit !== ledSuit) continue;
    return suit;
  }
  return null;
}

let pendingBelotPlay = null;

function playCardMaybeBelot(cardStr, m) {
  const { rank, suit } = parseCard(cardStr);
  const belotSuit = findBelotSuit(m.yourHand, m);
  if (belotSuit === suit && (rank === 'K' || rank === 'Q')) {
    pendingBelotPlay = cardStr;
    $('belotOverlay').classList.remove('hidden');
    return;
  }
  sendMsg({ type: 'play_card', card: cardStr });
}

// A plain click plays the card immediately. A real drag (mouse moves
// more than a few pixels before release) shows the card following the
// cursor for visual feedback, then plays it on release - landing spot
// barely matters since the server's authoritative trick-pile render
// replaces this element within a moment anyway ("система премества в
// центъра" for free, simply by virtue of the real trick display).
function attachClickOrDrag(el, cardStr, m) {
  el.addEventListener('mousedown', (downEv) => {
    if (downEv.button !== 0) return; // left click only
    const startX = downEv.clientX, startY = downEv.clientY;
    let dragging = false;
    let ghost = null;
    const rect = el.getBoundingClientRect();
    const offsetX = downEv.clientX - rect.left;
    const offsetY = downEv.clientY - rect.top;

    const onMove = (moveEv) => {
      const dx = moveEv.clientX - startX, dy = moveEv.clientY - startY;
      if (!dragging && (Math.abs(dx) > 6 || Math.abs(dy) > 6)) {
        dragging = true;
        ghost = el.cloneNode(true);
        ghost.classList.add('dragging-card');
        ghost.style.position = 'fixed';
        ghost.style.width = rect.width + 'px';
        ghost.style.height = rect.height + 'px';
        ghost.style.zIndex = '999';
        ghost.style.pointerEvents = 'none';
        ghost.style.transform = 'none';
        document.body.appendChild(ghost);
        el.style.opacity = '0.25';
      }
      if (dragging && ghost) {
        ghost.style.left = (moveEv.clientX - offsetX) + 'px';
        ghost.style.top = (moveEv.clientY - offsetY) + 'px';
      }
    };

    const onUp = () => {
      document.removeEventListener('mousemove', onMove);
      document.removeEventListener('mouseup', onUp);
      if (ghost) { ghost.remove(); }
      el.style.opacity = '';
      el.style.display = 'none';
      playCardMaybeBelot(cardStr, m);
    };

    document.addEventListener('mousemove', onMove);
    document.addEventListener('mouseup', onUp);
  });
}

$('belotYesBtn')?.addEventListener('click', () => {
  if (pendingBelotPlay) {
    sendMsg({ type: 'declare', kind: 'belot', value: 20, category: 'belot' });
    sendMsg({ type: 'play_card', card: pendingBelotPlay });
  }
  pendingBelotPlay = null;
  $('belotOverlay').classList.add('hidden');
});
$('belotNoBtn')?.addEventListener('click', () => {
  if (pendingBelotPlay) sendMsg({ type: 'play_card', card: pendingBelotPlay });
  pendingBelotPlay = null;
  $('belotOverlay').classList.add('hidden');
});

// Shows a small non-blocking prompt on a player's own first play of the
// hand, listing any sequence/carre combos their hand supports, with a
// button per combo to declare it out loud to the table.
let declaredCategoriesThisHand = new Set();

function renderAnnouncePrompt(m) {
  const panel = $('announcePrompt');
  const isMyFirstPlay = m.phase === 'playing' && m.turn === mySeat && m.yourHand.length === 8;
  const anns = m.myAnnounces || [];
  if (!isMyFirstPlay || anns.length === 0) {
    panel.classList.add('hidden');
    return;
  }
  panel.classList.remove('hidden');
  const btnContainer = $('announceButtons');
  btnContainer.innerHTML = '';
  anns.forEach(a => {
    const btn = document.createElement('button');
    const alreadyDeclared = declaredCategoriesThisHand.has(a.category);
    btn.textContent = t('announcePrefix', announceLabel(a));
    btn.disabled = alreadyDeclared;
    btn.addEventListener('click', () => {
      sendMsg({ type: 'declare', kind: a.kind, suit: a.suit, highRank: a.highRank, value: a.value, category: a.category });
      declaredCategoriesThisHand.add(a.category);
      renderAnnouncePrompt(m);
    });
    btnContainer.appendChild(btn);
  });
}

let ws = null;
let mySeat = -1;
let myName = '';
let myTeam = null; // 'cherry' | 'malina', set once chosen in the lobby
let lastState = null;

const SESSION_KEY = 'belotSession';
function saveSession(code, name) {
  try { localStorage.setItem(SESSION_KEY, JSON.stringify({ code, name })); } catch (e) { /* storage unavailable - not fatal */ }
}
function loadSession() {
  try {
    const raw = localStorage.getItem(SESSION_KEY);
    return raw ? JSON.parse(raw) : null;
  } catch (e) { return null; }
}
function clearSession() {
  try { localStorage.removeItem(SESSION_KEY); } catch (e) { /* ignore */ }
}

function teamColorClass(seat) { return teamOf(seat) === teamOf(mySeat) ? 'team-us' : 'team-them'; }
function teamDisplayName(teamIndex) {
  const name = teamNames[teamIndex];
  if (name === 'Отбор А') return t('teamADefault');
  if (name === 'Отбор Б') return t('teamBDefault');
  return name || (teamIndex === 0 ? t('teamADefault') : t('teamBDefault'));
}

function $(id) { return document.getElementById(id); }
function showScreen(id) {
  ['lobby', 'waiting', 'table'].forEach(s => $(s).classList.toggle('hidden', s !== id));
}

function parseCard(str) {
  // "10-hearts" / "J-clubs"
  const i = str.lastIndexOf('-');
  return { rank: str.slice(0, i), suit: str.slice(i + 1) };
}

function cardEl(cardStr, opts) {
  opts = opts || {};
  const { rank, suit } = parseCard(cardStr);
  const el = document.createElement('div');
  el.className = 'card' + (opts.extraClass ? ' ' + opts.extraClass : '');
  const img = document.createElement('img');
  img.src = `cards/${rank}${SUIT_LETTER[suit]}.svg`;
  img.alt = `${rank} ${SUIT_SYMBOL[suit]}`;
  img.draggable = false;
  el.appendChild(img);
  if (opts.disabled) el.classList.add('disabled');
  if (opts.onClick) el.addEventListener('click', opts.onClick);
  return el;
}

function cardBackEl() {
  const el = document.createElement('div');
  el.className = 'card-back';
  const img = document.createElement('img');
  img.src = 'cards/back.png';
  img.alt = '';
  img.draggable = false;
  el.appendChild(img);
  return el;
}

// Fans a set of face-down card-backs the same way the own hand fans
// out - slight rotation + overlap - either horizontally (top seat) or
// vertically (left/right seats, stacked top to bottom).
function fanCardBacks(container, count, axis) {
  container.innerHTML = '';
  const mid = (count - 1) / 2;
  const ANGLE_STEP = 6;
  for (let i = 0; i < count; i++) {
    const el = cardBackEl();
    const angle = (i - mid) * ANGLE_STEP * (axis === 'vertical' ? -1 : 1);
    el.style.setProperty('--angle', angle + 'deg');
    if (i > 0) {
      if (axis === 'vertical') el.style.marginTop = '-58px';
      else el.style.marginLeft = '-40px';
    }
    container.appendChild(el);
  }
}

// Cards land close to the center but offset toward whichever relative
// direction (bottom/left/top/right) their player sits at, so a card
// ends up positioned roughly facing the player who threw it - with a
// little corner overlap between neighboring cards, not stacked
// directly on top of each other like a single messy pile.
const TRICK_SLOT_OFFSET = {
  0: { dx: 0, dy: 130 }, // bottom (me)
  1: { dx: 130, dy: 0 }, // pos 1 renders on the RIGHT (seat3/.seat-right)
  2: { dx: 0, dy: -130 }, // top
  3: { dx: -130, dy: 0 }, // pos 3 renders on the LEFT (seat1/.seat-left)
};

// Random resting rotation per player for the current trick - a few
// degrees, randomized fresh each trick, but stable across re-renders
// of the same (already-settled) card so it doesn't jitter.
let trickCardRotations = {};

function pileCardTransform(playerSeat) {
  const pos = seatDelta(playerSeat);
  const off = TRICK_SLOT_OFFSET[pos] || { dx: 0, dy: 0 };
  if (trickCardRotations[playerSeat] === undefined) {
    trickCardRotations[playerSeat] = Math.random() * 30 - 15; // -15..+15deg
  }
  const rot = trickCardRotations[playerSeat];
  return { rot, transform: `translate(-50%, -50%) translate(${off.dx}px, ${off.dy}px) rotate(${rot}deg)` };
}

let lastTrickCardCount = 0;

// Renders a set of {player, card} entries near the center of the
// trick area, each positioned toward the side of whoever played it,
// slightly overlapping corners with its neighbors - not stacked
// directly on top of each other. Cards that are NEW since the
// previous render (i.e. just played) fly in from the playing seat's
// position on screen, face up, spinning a full turn as they travel -
// already-settled cards from earlier in the same trick are just
// redrawn in place, no re-animation. Because each viewer computes the
// seat's screen position independently (seatDelta is relative to
// their own mySeat), the same play looks like it's coming from a
// different direction on every player's own screen.
function renderTrickPile(cards) {
  const trickArea = $('trickArea');
  cards = cards || [];
  const prevCount = cards.length >= lastTrickCardCount ? lastTrickCardCount : 0;
  trickArea.innerHTML = '';

  cards.forEach((pc, i) => {
    if (i < prevCount) {
      const { transform } = pileCardTransform(pc.player);
      const el = cardEl(pc.card, { extraClass: 'card-static trick-pile-card' });
      el.dataset.player = pc.player;
      el.style.transform = transform;
      el.style.zIndex = i;
      trickArea.appendChild(el);
    } else {
      animateCardIntoPile(trickArea, pc, i);
    }
  });
  lastTrickCardCount = cards.length;
}

const CARD_W = 120, CARD_H = 172; // matches --card-w / --card-h

function animateCardIntoPile(trickArea, pc, i) {
  const { rot, transform } = pileCardTransform(pc.player);

  const finalEl = cardEl(pc.card, { extraClass: 'card-static trick-pile-card' });
  finalEl.dataset.player = pc.player;
  finalEl.style.transform = transform;
  finalEl.style.zIndex = i;
  finalEl.style.visibility = 'hidden';
  trickArea.appendChild(finalEl);

  const endRect = finalEl.getBoundingClientRect();
  const endX = endRect.left + endRect.width / 2;
  const endY = endRect.top + endRect.height / 2;

  const startEl = seatElementFor(pc.player);
  const startRect = startEl ? startEl.getBoundingClientRect() : endRect;
  const startX = startRect.left + startRect.width / 2;
  const startY = startRect.top + startRect.height / 2;

  const flying = cardEl(pc.card, {});
  flying.classList.add('flying-card');
  flying.style.left = (startX - CARD_W / 2) + 'px';
  flying.style.top = (startY - CARD_H / 2) + 'px';
  flying.style.transform = 'rotate(0deg)';
  document.body.appendChild(flying);
  void flying.offsetWidth;
  requestAnimationFrame(() => {
    flying.style.transition = 'left 0.3s ease-out, top 0.3s ease-out, transform 0.3s ease-out';
    flying.style.left = (endX - CARD_W / 2) + 'px';
    flying.style.top = (endY - CARD_H / 2) + 'px';
    // A full spin while traveling, landing exactly on the pile's own
    // (randomized) resting rotation.
    flying.style.transform = `rotate(${360 + rot}deg)`;
  });
  setTimeout(() => {
    flying.remove();
    finalEl.style.visibility = 'visible';
  }, 300);
}

let reconnectAttempts = 0;
let reconnectTimer = null;
let intentionalDisconnect = false;

function connect() {
  const proto = location.protocol === 'https:' ? 'wss' : 'ws';
  ws = new WebSocket(`${proto}://${location.host}/ws`);
  ws.onmessage = (ev) => handleMessage(JSON.parse(ev.data));
  ws.onclose = () => {
    if (intentionalDisconnect) return;
    const session = loadSession();
    if (!session) return; // never actually joined anything - nothing to resume
    scheduleReconnect(session);
  };
}

function scheduleReconnect(session) {
  reconnectAttempts++;
  $('connStatus').textContent = t('reconnecting');
  $('connStatus').classList.remove('hidden');
  const delay = Math.min(1000 * reconnectAttempts, 5000);
  clearTimeout(reconnectTimer);
  reconnectTimer = setTimeout(() => {
    connect();
    ws.onopen = () => {
      reconnectAttempts = 0;
      $('connStatus').classList.add('hidden');
      sendMsg({ type: 'join_room', name: session.name, code: session.code });
    };
  }, delay);
}

function handleMessage(m) {
  if (showingTrickResult && (m.type === 'hand_result' || m.type === 'match_over')) {
    pendingMessages.push(m);
    return;
  }
  switch (m.type) {
    case 'room_state': return onRoomState(m);
    case 'prompt_team_name': return onPromptTeamName(m);
    case 'game_state': return onGameState(m);
    case 'trick_complete': return onTrickComplete(m);
    case 'declared': return onDeclared(m);
    case 'ready_status': return onReadyStatus(m);
    case 'hand_result': return onHandResult(m);
    case 'match_over': return onMatchOver(m);
    case 'error': return onError(m);
  }
}

function seatElementFor(seat) {
  if (seat === mySeat) return document.querySelector('.my-area');
  const pos = seatDelta(seat);
  const elId = pos === 1 ? 'seat3' : pos === 2 ? 'seat2' : 'seat1';
  return $(elId);
}

let showingTrickResult = false;
let pendingGameState = null;
let pendingMessages = [];
let declaredLabels = { 0: [], 1: [], 2: [], 3: [] };
let declaredHandNumber = null;
let scoreRows = { 0: [], 1: [] };
let lastSeenHandNumber = null;
let handHistory = []; // [{cards:[{player,card}], winnerTeam}] in play order, this hand only
let iClickedReady = false;

function pushScoreRow(teamIdx, value) {
  if (!scoreRows[teamIdx]) scoreRows[teamIdx] = [];
  scoreRows[teamIdx].push(value);
}

function renderScoreboard() {
  const usIdx = teamOf(mySeat);
  const themIdx = usIdx === 0 ? 1 : 0;
  const usRows = scoreRows[usIdx] || [];
  const themRows = scoreRows[themIdx] || [];
  $('scoreboardHeaderUs').textContent = teamDisplayName(usIdx);
  $('scoreboardHeaderThem').textContent = teamDisplayName(themIdx);
  $('scoreRowsUs').innerHTML = usRows.map(v => `<div>${v}</div>`).join('');
  $('scoreRowsThem').innerHTML = themRows.map(v => `<div>${v}</div>`).join('');
  $('scoreTotalUs').textContent = usRows.reduce((a, b) => a + b, 0);
  $('scoreTotalThem').textContent = themRows.reduce((a, b) => a + b, 0);
}

function onDeclared(m) {
  if (!declaredLabels[m.player]) declaredLabels[m.player] = [];
  declaredLabels[m.player].push({ kind: m.kind, suit: m.suit, highRank: m.highRank, value: m.value });
  renderDeclaredLabels();
}

function renderDeclaredLabels() {
  for (let seat = 0; seat < 4; seat++) {
    const text = (declaredLabels[seat] || []).map(announceLabel).join(' · ');
    const el = seat === mySeat ? $('myDeclaredLabel') : seatElementFor(seat)?.querySelector('.declared-label');
    if (el) el.textContent = text;
  }
}

function onTrickComplete(m) {
  handHistory.push({ cards: m.cards, winnerTeam: teamOf(m.winner) });
  showingTrickResult = true;
  renderTrickPile(m.cards);

  setTimeout(() => {
    flyTrickToWinner(m.winner);
    setTimeout(() => {
      $('trickArea').innerHTML = '';
      lastTrickCardCount = 0;
      trickCardRotations = {};
      showingTrickResult = false;
      if (pendingGameState) {
        const s = pendingGameState;
        pendingGameState = null;
        renderGameState(s);
      }
      const queued = pendingMessages;
      pendingMessages = [];
      queued.forEach(handleMessage);
    }, 650);
  }, 3000);
}

function flyTrickToWinner(winnerSeat) {
  const targetEl = seatElementFor(winnerSeat);
  if (!targetEl) return;
  const targetRect = targetEl.getBoundingClientRect();
  const targetX = targetRect.left + targetRect.width / 2;
  const targetY = targetRect.top + targetRect.height / 2;
  const cardW = 120, cardH = 172;

  $('trickArea').querySelectorAll('.card').forEach((el, i) => {
    const rect = el.getBoundingClientRect();
    const startX = rect.left + rect.width / 2;
    const startY = rect.top + rect.height / 2;
    el.remove();

    const flying = el.cloneNode(true);
    flying.classList.add('flying-card');
    flying.classList.remove('trick-pile-card');
    flying.style.left = (startX - cardW / 2) + 'px';
    flying.style.top = (startY - cardH / 2) + 'px';
    flying.style.transform = 'rotate(0deg)';
    document.body.appendChild(flying);
    void flying.offsetWidth;
    requestAnimationFrame(() => {
      flying.style.left = (targetX - cardW / 2) + 'px';
      flying.style.top = (targetY - cardH / 2) + 'px';
      flying.style.opacity = '0.15';
      flying.style.transform = 'scale(0.4) rotate(' + (i * 20 - 20) + 'deg)';
    });
    setTimeout(() => flying.remove(), 650);
  });
}

// Maps the server's language-neutral error codes to a translated message.
// Unrecognized codes (e.g. raw Go error text from rare defensive paths)
// fall back to showing the code as-is rather than crashing the UI.
const ERROR_CODE_KEYS = {
  room_not_found: 'errorRoomNotFound',
  room_full: 'errorRoomFull',
  teams_incomplete_before_start: 'errorTeamsIncomplete',
  invalid_team: 'errorInvalidTeam',
  team_full: 'errorTeamFull',
};

function onError(m) {
  const el = $('lobbyError');
  const key = ERROR_CODE_KEYS[m.message];
  if (el) { el.textContent = key ? t(key) : m.message; }
  console.warn('server error:', m.message);
  if (m.message === 'room_not_found') {
    clearSession();
    showScreen('lobby');
  }
}

let teamNames = ['Отбор А', 'Отбор Б'];

let currentRoomCode = '';
function renderRoomCodeLabel() {
  $('roomCodeLabel').innerHTML = t('roomCodeLabel') + ' <span id="roomCode" class="room-code"></span>';
  $('roomCode').textContent = currentRoomCode;
}

function onRoomState(m) {
  currentRoomCode = m.code;
  renderRoomCodeLabel();
  saveSession(m.code, myName);
  const mySlot = m.yourSlot;
  const list = $('playerList');
  list.innerHTML = '';
  let joined = 0;
  m.players.forEach((p, i) => {
    const li = document.createElement('li');
    if (p) { li.textContent = `${t('seatLabel', i + 1)}: ${p.name}`; joined++; }
    else { li.textContent = t('seatEmpty', i + 1); }
    list.appendChild(li);
  });
  showScreen('waiting');

  if (Array.isArray(m.teamNames) && m.teamNames.length === 2) teamNames = m.teamNames;

  const teamSelect = $('teamSelect');
  teamSelect.classList.toggle('hidden', joined < 4);

  const cherryNames = [];
  const malinaNames = [];
  let cherryCount = 0, malinaCount = 0;
  m.players.forEach((p, i) => {
    if (!p) return;
    if (p.team === 'cherry') { cherryNames.push(p.name); cherryCount++; if (i === mySlot) myTeam = 'cherry'; }
    else if (p.team === 'malina') { malinaNames.push(p.name); malinaCount++; if (i === mySlot) myTeam = 'malina'; }
    else if (i === mySlot) { myTeam = null; }
  });

  $('cherryList').innerHTML = cherryNames.map(n => `<li>${n}</li>`).join('');
  $('malinaList').innerHTML = malinaNames.map(n => `<li>${n}</li>`).join('');
  $('chooseCherryBtn').classList.toggle('taken', cherryCount >= 2 && myTeam !== 'cherry');
  $('chooseCherryBtn').classList.toggle('mine', myTeam === 'cherry');
  $('chooseMalinaBtn').classList.toggle('taken', malinaCount >= 2 && myTeam !== 'malina');
  $('chooseMalinaBtn').classList.toggle('mine', myTeam === 'malina');

  const cherryPrompted = Array.isArray(m.teamNamePrompted) && m.teamNamePrompted[0];
  const malinaPrompted = Array.isArray(m.teamNamePrompted) && m.teamNamePrompted[1];
  const cherryNamed = cherryPrompted && teamNames[0] !== 'Отбор А';
  const malinaNamed = malinaPrompted && teamNames[1] !== 'Отбор Б';
  $('cherryNameStatus').textContent = cherryNamed ? t('namedAs', teamNames[0]) : (cherryPrompted ? t('choosingName') : '');
  $('malinaNameStatus').textContent = malinaNamed ? t('namedAs', teamNames[1]) : (malinaPrompted ? t('choosingName') : '');

  const teamsReady = cherryCount === 2 && malinaCount === 2 && cherryNamed && malinaNamed;
  $('startBtn').classList.toggle('hidden', joined < 4 || !teamsReady);
}

function onPromptTeamName(m) {
  $('teamNameInput').value = '';
  $('teamNameOverlay').classList.remove('hidden');
  $('teamNameInput').focus();
}

$('teamNameDoneBtn')?.addEventListener('click', () => {
  const name = $('teamNameInput').value.trim();
  if (!name) return;
  sendMsg({ type: 'set_team_name', teamName: name });
  $('teamNameOverlay').classList.add('hidden');
});

function seatDelta(seat) {
  // Abstract relative offset from "me" (0=me/bottom), NOT a direct
  // left/right label - see seatElementFor for the actual screen-side
  // mapping (pos 1 -> right, pos 3 -> left, after the counter-
  // clockwise turn-direction fix swapped which physical side each
  // uses).
  return (seat - mySeat + 4) % 4;
}

function onGameState(m) {
  lastState = m;
  mySeat = m.yourSeat;
  if (showingTrickResult) {
    pendingGameState = m;
    return;
  }
  renderGameState(m);
}

// Builds the display text for a sequence/carre/belot announce from the
// server's language-neutral fields (kind/suit/highRank/value), so it
// renders in whichever language the viewer currently has selected -
// unlike a pre-rendered label, which would be stuck in whoever declared
// it's language.
function announceLabel(a) {
  const rank = (t('rankDisplay')[a.highRank]) || a.highRank;
  const suitName = a.suit ? (t('suitNames')[a.suit] || '') : '';
  let text = '';
  switch (a.kind) {
    case 'tierce': text = t('annTierce', rank, suitName); break;
    case 'fifty': text = t('annFifty', rank, suitName); break;
    case 'hundred': text = t('annHundred', rank, suitName); break;
    case 'carreJ': text = t('annCarreJ'); break;
    case 'carreNine': text = t('annCarreNine'); break;
    case 'carreOther': text = t('annCarreOther', rank); break;
    case 'belot': text = t('annBelot'); break;
  }
  return `${text} (${a.value})`;
}

function contractShortLabel(type, suit) {
  if (type === 'suit') {
    const s = t('suitNames')[suit] || '';
    return s.charAt(0).toUpperCase() + s.slice(1);
  }
  if (type === 'notrump') return t('noTrump');
  if (type === 'alltrump') return t('allTrump');
  return '';
}

let bidHistory = [];
let bidHistoryHandNumber = null;
let prevBidContract = null;
let prevBidTurn = null;

function playerNameFor(m, seat) {
  return (m.players && m.players[seat]) ? m.players[seat] : t('seatLabel', seat + 1);
}

function updateBidHistory(m) {
  if (bidHistoryHandNumber !== m.handNumber) {
    bidHistoryHandNumber = m.handNumber;
    bidHistory = [];
    prevBidContract = null;
    prevBidTurn = null;
  }

  if (m.contract && prevBidTurn !== null) {
    const c = m.contract;
    const prev = prevBidContract;
    const isNewCall = !prev || prev.type !== c.type || prev.suit !== c.suit;
    if (isNewCall) {
      bidHistory.push(`${contractShortLabel(c.type, c.suit)} (${playerNameFor(m, prevBidTurn)})`);
    } else if (c.contra && !(prev && prev.contra)) {
      bidHistory.push(t('contraLine', playerNameFor(m, prevBidTurn)));
    } else if (c.reconto && !(prev && prev.reconto)) {
      bidHistory.push(t('recontoLine', playerNameFor(m, prevBidTurn)));
    }
  }
  prevBidContract = m.contract ? { type: m.contract.type, suit: m.contract.suit, contra: m.contract.contra, reconto: m.contract.reconto } : null;
  prevBidTurn = m.turn;

  renderBidHistory(m);
}

function renderBidHistory(m) {
  let html = bidHistory.map(line => `<div class="bid-history-line">${line}</div>`).join('');
  if (m.contract && m.phase !== 'bidding') {
    const callerName = playerNameFor(m, m.contract.callerId);
    let finalText = `${contractShortLabel(m.contract.type, m.contract.suit).toUpperCase()} (${callerName})`;
    if (m.contract.reconto) finalText += t('withReconto');
    else if (m.contract.contra) finalText += t('withContra');
    html += `<div class="bid-history-final">${finalText}</div>`;
  }
  $('bidHistory').innerHTML = html;
}

// Shows the strongest->weakest rank order for the current contract, below
// the scoreboard. Всичко коз / без коз get one universal line since every
// suit ranks the same way; a suit contract gets two lines - the trump suit
// (all-trump order) and everyone else (no-trump order).
function renderCardStrengthHint(m) {
  const hint = $('cardStrengthHint');
  if (!m.contract || m.phase === 'bidding') {
    hint.classList.add('hidden');
    return;
  }
  hint.classList.remove('hidden');
  $('cardStrengthTitle').textContent = t('cardStrengthTitle');
  const line2 = $('cardStrengthLine2');
  if (m.contract.type === 'suit') {
    const suitName = t('suitNamesCap')[m.contract.suit] || '';
    $('cardStrengthLine1').textContent = `${suitName} - ${TRUMP_STRENGTH_ORDER.join(' ')}`;
    line2.textContent = `${t('cardStrengthOther')} - ${NORMAL_STRENGTH_ORDER.join(' ')}`;
    line2.classList.remove('hidden');
  } else {
    const order = m.contract.type === 'alltrump' ? TRUMP_STRENGTH_ORDER : NORMAL_STRENGTH_ORDER;
    $('cardStrengthLine1').textContent = order.join(' ');
    line2.classList.add('hidden');
  }
}

function renderGameState(m) {
  showScreen('table');
  if (Array.isArray(m.teamNames) && m.teamNames.length === 2) teamNames = m.teamNames;

  if (declaredHandNumber !== m.handNumber) {
    declaredHandNumber = m.handNumber;
    declaredLabels = { 0: [], 1: [], 2: [], 3: [] };
    declaredCategoriesThisHand = new Set();
    handHistory = [];
    iClickedReady = false;
    $('handRecapOverlay').classList.add('hidden');
  }
  if (lastSeenHandNumber !== 1 && m.handNumber === 1) {
    scoreRows = { 0: [], 1: [] };
  }
  lastSeenHandNumber = m.handNumber;
  renderDeclaredLabels();
  renderScoreboard();

  updateBidHistory(m);
  renderCardStrengthHint(m);

  // Other seats: name (colored by team) + face-down full-size cards by count.
  for (let seat = 0; seat < 4; seat++) {
    if (seat === mySeat) continue;
    const pos = seatDelta(seat);
    const elId = pos === 1 ? 'seat3' : pos === 2 ? 'seat2' : 'seat1';
    const el = $(elId);
    el.classList.toggle('active', m.turn === seat);
    const nameEl = el.querySelector('.seat-name');
    nameEl.textContent = (m.players && m.players[seat] ? m.players[seat] : t('seatLabel', seat + 1)) + (m.dealer === seat ? t('dealerSuffixParen') : '');
    nameEl.className = 'seat-name ' + teamColorClass(seat);
    const cardsEl = el.querySelector('.seat-cards');
    cardsEl.innerHTML = '';
    const count = m.handCounts[seat] || 0;
    const axis = (pos === 1 || pos === 3) ? 'vertical' : 'horizontal';
    fanCardBacks(cardsEl, count, axis);
  }

  const myDisplayName = (m.players && m.players[mySeat]) ? m.players[mySeat] : t('seatLabel', mySeat + 1);
  const mySeatNameEl = $('mySeatName');
  mySeatNameEl.textContent = t('youLabel', myDisplayName) + (m.dealer === mySeat ? t('dealerSuffixDash') : '');
  mySeatNameEl.className = 'seat-name ' + teamColorClass(mySeat);
  mySeatNameEl.classList.toggle('active-turn', m.turn === mySeat);

  // Current trick - plus/cross layout, each card pointing from its
  // player's seat toward the center.
  renderTrickPile(m.currentTrick);

  // My hand - sorted by suit (♠♥♦♣) and strength, fanned out with each
  // card overlapping the next (~1/3 of the top-left corner peeking out).
  const handEl = $('myHand');
  handEl.innerHTML = '';
  const legal = new Set(m.legalMoves || []);
  const canPlay = m.phase === 'playing' && m.turn === mySeat;
  const sortedHand = sortHand(m.yourHand, m);
  const n = sortedHand.length;
  const mid = (n - 1) / 2;
  const ANGLE_STEP = 6; // degrees of fan spread per card
  sortedHand.forEach((cardStr, i) => {
    const clickable = canPlay && legal.has(cardStr);
    // Only visually dim a card when it's genuinely your turn to play
    // and this particular card is illegal - not just because it isn't
    // your turn at all (e.g. during bidding, or someone else's turn),
    // which would otherwise wash out your whole hand.
    const dim = canPlay && !clickable;
    const el = cardEl(cardStr, { disabled: dim });
    if (clickable) attachClickOrDrag(el, cardStr, m);
    const angle = (i - mid) * ANGLE_STEP;
    el.style.setProperty('--angle', angle + 'deg');
    el.style.zIndex = i;
    if (i > 0) el.style.marginLeft = '-80px';
    handEl.appendChild(el);
  });

  renderAnnouncePrompt(m);

  // Bidding panel.
  const biddingPanel = $('biddingPanel');
  const myTurnToBid = m.phase === 'bidding' && m.turn === mySeat;
  biddingPanel.classList.toggle('hidden', !myTurnToBid);
  if (myTurnToBid) {
    const hasContract = !!m.contract;
    const currentRank = hasContract ? callRank(m.contract.type, m.contract.suit) : -1;
    $('biddingTurnLabel').textContent = hasContract
      ? t('bidPromptHasContract')
      : t('bidPromptNoContract');
    biddingPanel.querySelectorAll('[data-suit]').forEach(b => {
      b.classList.toggle('hidden', callRank('suit', b.dataset.suit) <= currentRank);
    });
    $('noTrumpBtn').classList.toggle('hidden', callRank('notrump') <= currentRank);
    $('allTrumpBtn').classList.toggle('hidden', callRank('alltrump') <= currentRank);
    const opponentContract = hasContract && teamOf(m.contract.callerId) !== teamOf(mySeat);
    $('contraBtn').classList.toggle('hidden', !(opponentContract && !m.contract.contra));
    const ownContractContra = hasContract && teamOf(m.contract.callerId) === teamOf(mySeat) && m.contract.contra && !m.contract.reconto;
    $('recontoBtn').classList.toggle('hidden', !ownContractContra);
  }
}

// Standard Bulgarian belot auction order, weakest to strongest:
// ♣ спатия < ♦ каро < ♥ купа < ♠ пика < без коз < всичко коз.
const SUIT_BID_ORDER = ['clubs', 'diamonds', 'hearts', 'spades'];
function callRank(type, suit) {
  if (type === 'suit') return SUIT_BID_ORDER.indexOf(suit);
  if (type === 'notrump') return 4;
  if (type === 'alltrump') return 5;
  return -1;
}

function teamOf(seat) { return seat % 2; }

let lastHandPoints = { 0: 0, 1: 0 };
let lastHandWinnerTeam = null;

function onHandResult(m) {
  const winnerPoints = Math.floor(m.awardedPoints / 10);
  const otherTeam = m.awardedTeam === 0 ? 1 : 0;
  const otherPoints = m.defenderPoints > 0 ? Math.floor(m.defenderPoints / 10) : 0;
  pushScoreRow(m.awardedTeam, winnerPoints);
  pushScoreRow(otherTeam, otherPoints);
  lastHandPoints = { [m.awardedTeam]: winnerPoints, [otherTeam]: otherPoints };
  lastHandWinnerTeam = m.awardedTeam;
  renderScoreboard();
  showHandRecap();
}

const CARD_ASPECT = 120 / 172; // matches --card-w / --card-h

function layoutRecapTeamRow(containerEl, cards) {
  containerEl.innerHTML = '';
  const n = cards.length;
  if (n === 0) return;

  const rowH = containerEl.clientHeight;
  const rowW = containerEl.clientWidth;
  const cardH = Math.max(30, rowH);
  const cardW = cardH * CARD_ASPECT;

  // Minimum visible sliver of an overlapped card must still show its
  // top-left rank+suit corner - roughly a third of the card's width.
  const minVisible = cardW * 0.34;
  let overlap = 0;
  if (n > 1) {
    const naturalWidth = n * cardW;
    if (naturalWidth > rowW) {
      overlap = cardW - (rowW - cardW) / (n - 1);
      const maxOverlap = cardW - minVisible;
      overlap = Math.min(Math.max(overlap, 0), maxOverlap);
    }
  }

  cards.forEach((pc, i) => {
    const el = cardEl(pc.card, { extraClass: 'card-static recap-card' });
    el.style.width = cardW + 'px';
    el.style.height = cardH + 'px';
    el.style.zIndex = i;
    if (i > 0) el.style.marginLeft = -overlap + 'px';
    containerEl.appendChild(el);
  });
}

function layoutRecap() {
  if ($('handRecapOverlay').classList.contains('hidden')) return;
  const usIdx = teamOf(mySeat);
  const usCards = [];
  const themCards = [];
  handHistory.forEach(trick => {
    trick.cards.forEach(pc => {
      (trick.winnerTeam === usIdx ? usCards : themCards).push(pc);
    });
  });
  layoutRecapTeamRow($('recapRowUs'), usCards);
  layoutRecapTeamRow($('recapRowThem'), themCards);
}

window.addEventListener('resize', layoutRecap);

function showHandRecap() {
  const usIdx = teamOf(mySeat);
  const themIdx = usIdx === 0 ? 1 : 0;
  const usEl = $('recapLabelUs');
  const themEl = $('recapLabelThem');
  usEl.textContent = `${teamDisplayName(usIdx)} - ${lastHandPoints[usIdx] || 0}`;
  themEl.textContent = `${teamDisplayName(themIdx)} - ${lastHandPoints[themIdx] || 0}`;
  usEl.classList.toggle('recap-label-winner', lastHandWinnerTeam === usIdx);
  themEl.classList.toggle('recap-label-winner', lastHandWinnerTeam === themIdx);

  iClickedReady = false;
  $('recapDoneBtn').classList.remove('hidden');
  $('recapWaiting').textContent = '';
  $('handRecapOverlay').classList.remove('hidden');
  requestAnimationFrame(layoutRecap);
}

function onReadyStatus(m) {
  if (iClickedReady) {
    $('recapWaiting').textContent = t('waitingPlayersCount', m.ready, m.total);
  }
}

$('recapDoneBtn')?.addEventListener('click', () => {
  iClickedReady = true;
  sendMsg({ type: 'ready_for_next_hand' });
  $('recapDoneBtn').classList.add('hidden');
  $('recapWaiting').textContent = t('waitingPlayers');
});

$('newGameBtn')?.addEventListener('click', () => {
  $('resultOverlay').classList.add('hidden');
  sendMsg({ type: 'start_game' });
});

function onMatchOver(m) {
  $('resultTitle').textContent = t('gameOver');
  const bodyEl = $('resultBody');
  bodyEl.textContent = t('teamWinsMatch', teamDisplayName(m.winner));
  bodyEl.style.color = '#ff4d4d';
  bodyEl.style.fontWeight = 'bold';
  $('resultCloseBtn').classList.add('hidden');
  $('newGameBtn').classList.remove('hidden');
  $('resultOverlay').classList.remove('hidden');
}

function sendMsg(obj) {
  if (ws && ws.readyState === WebSocket.OPEN) ws.send(JSON.stringify(obj));
}

$('createBtn').addEventListener('click', () => {
  const name = $('nameInput').value.trim();
  if (!name) { $('lobbyError').textContent = t('enterYourName'); return; }
  myName = name;
  connect();
  ws.onopen = () => sendMsg({ type: 'create_room', name: myName, targetScore: 151 });
});

$('joinBtn').addEventListener('click', () => {
  const name = $('nameInput').value.trim();
  if (!name) { $('lobbyError').textContent = t('enterYourName'); return; }
  myName = name;
  const code = $('codeInput').value.trim().toUpperCase();
  if (!code) { $('lobbyError').textContent = t('enterRoomCode'); return; }
  connect();
  ws.onopen = () => sendMsg({ type: 'join_room', name: myName, code });
});

$('leaveRoomBtn')?.addEventListener('click', () => {
  if (!confirm(t('confirmLeaveRoom'))) return;
  intentionalDisconnect = true;
  clearSession();
  location.reload();
});

// On page load, if a previous session was saved (created/joined a room
// earlier - possibly a while ago), try to silently rejoin it. If the
// room is gone (server restarted since) or something else goes wrong,
// onError already falls back to the normal lobby screen and clears the
// stale session. This is what makes a refresh (or reopening the tab
// the next day, as long as the server itself hasn't restarted) resume
// right where the game was, instead of dumping you back at square one.
(function tryAutoRejoin() {
  const session = loadSession();
  if (!session || !session.code || !session.name) return;
  myName = session.name;
  const nameInput = $('nameInput');
  if (nameInput) nameInput.value = session.name;
  const errEl = $('lobbyError');
  if (errEl) errEl.textContent = t('connectingPrevious');
  connect();
  ws.onopen = () => sendMsg({ type: 'join_room', name: session.name, code: session.code });
})();

$('startBtn').addEventListener('click', () => sendMsg({ type: 'start_game' }));
$('chooseCherryBtn').addEventListener('click', () => sendMsg({ type: 'choose_team', team: 'cherry' }));
$('chooseMalinaBtn').addEventListener('click', () => sendMsg({ type: 'choose_team', team: 'malina' }));

document.querySelectorAll('[data-suit]').forEach(btn => {
  btn.addEventListener('click', () => sendMsg({ type: 'bid', action: 'suit', suit: btn.dataset.suit }));
});
$('noTrumpBtn').addEventListener('click', () => sendMsg({ type: 'bid', action: 'notrump' }));
$('allTrumpBtn').addEventListener('click', () => sendMsg({ type: 'bid', action: 'alltrump' }));
$('passBtn').addEventListener('click', () => sendMsg({ type: 'bid', action: 'pass' }));
$('contraBtn').addEventListener('click', () => sendMsg({ type: 'bid', action: 'contra' }));
$('recontoBtn').addEventListener('click', () => sendMsg({ type: 'bid', action: 'reconto' }));

// Apply the saved (or default) language to every static label right
// away, and keep <html lang="..."> in sync.
document.documentElement.lang = currentLang;
applyStaticTranslations();
