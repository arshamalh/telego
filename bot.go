package telebot

import (
	"errors"
	"os"

	tba "github.com/SakoDroid/telebot/TBA"
	cfg "github.com/SakoDroid/telebot/configs"
	logger "github.com/SakoDroid/telebot/logger"
	objs "github.com/SakoDroid/telebot/objects"
)

type Bot struct {
	botCfg             *cfg.BotConfigs
	apiInterface       *tba.BotAPIInterface
	updateChannel      *chan *objs.Update
	pollUpdateChannel  *chan *objs.Poll
	pollRoutineChannel *chan bool
}

/*Starts the bot. If the bot has already been started it returns an error.*/
func (bot *Bot) Run() error {
	logger.InitTheLogger(bot.botCfg)
	go bot.startPollUpdateRoutine()
	return bot.apiInterface.StartUpdateRoutine()
}

/*Returns the channel which new updates received from api server are pushed into.*/
func (bot *Bot) GetUpdateChannel() *chan *objs.Update {
	return bot.updateChannel
}

/*Send a text message to a chat (not channel, use SendMessageToChannel method for sending messages to channles) and returns the sent message on success
If you want to ignore "parseMode" pass empty string. To ignore replyTo pass 0.*/
func (bot *Bot) SendMessage(chatId int, text, parseMode string, replyTo int, silent bool) (*objs.SendMethodsResult, error) {
	return bot.apiInterface.SendMessage(chatId, "", text, parseMode, nil, false, silent, false, replyTo, nil)
}

/*Send a text message to a channel and returns the sent message on success
If you want to ignore "parseMode" pass empty string. To ignore replyTo pass 0.*/
func (bot *Bot) SendMesssageToChannel(chatId, text, parseMode string, replyTo int, silent bool) (*objs.SendMethodsResult, error) {
	return bot.apiInterface.SendMessage(0, chatId, text, parseMode, nil, false, silent, false, replyTo, nil)
}

/*Returns a MessageForwarder which has several methods for forwarding a message*/
func (bot *Bot) ForwardMessage(messageId int, disableNotif bool) *MessageForwarder {
	return &MessageForwarder{bot: bot, messageId: messageId, disableNotif: disableNotif}
}

/*Returns a MessageCopier which has several methods for copying a message*/
func (bot *Bot) CopyMessage(messageId int, disableNotif bool) *MessageCopier {
	return &MessageCopier{bot: bot, messageId: messageId, disableNotif: disableNotif}
}

/*Returns a PhotoSender which has several methods for sending a photo. This method is only used for sending a photo to all types of chat except channels. To send a photo to a channel use "SendPhotoToChannel" method.
To ignore int arguments pass 0 and to ignore string arguments pass empty string ("")*/
func (bot *Bot) SendPhoto(chatId, replyTo int, caption, parseMode string) *PhotoSender {
	return &PhotoSender{bot: bot, chatIdInt: chatId, replyTo: replyTo, caption: caption, parseMode: parseMode}
}

/*Returns a PhotoSender which has several methods for sending a photo. This method is only used for sending a photo to a channels.
To ignore int arguments pass 0 and to ignore string arguments pass empty string ("")
*/
func (bot *Bot) SendPhotoToChannel(chatId string, replyTo int, caption, parseMode string) *PhotoSender {
	return &PhotoSender{bot: bot, chatIdInt: 0, chatidString: chatId, replyTo: replyTo, caption: caption, parseMode: parseMode}
}

/*Returns a VideoSender which has several methods for sending a video. This method is only used for sending a video to all types of chat except channels. To send a video to a channel use "SendVideoToChannel" method.
To ignore int arguments pass 0 and to ignore string arguments pass empty string ("")

---------------------------------

Official telegram doc :

Use this method to send video files, Telegram clients support mp4 videos (other formats may be sent as Document). On success, the sent Message is returned. Bots can currently send video files of up to 50 MB in size, this limit may be changed in the future.*/
func (bot *Bot) SendVideo(chatId int, replyTo int, caption, parseMode string) *VideoSender {
	return &VideoSender{bot: bot, chatIdInt: chatId, chatidString: "", replyTo: replyTo, caption: caption, parseMode: parseMode}
}

/*Returns a VideoSender which has several methods for sending a video. This method is only used for sending a video to a channels.
To ignore int arguments pass 0 and to ignore string arguments pass empty string ("")

---------------------------------

Official telegram doc :

Use this method to send video files, Telegram clients support mp4 videos (other formats may be sent as Document). On success, the sent Message is returned. Bots can currently send video files of up to 50 MB in size, this limit may be changed in the future.*/
func (bot *Bot) SendVideoToChannel(chatId string, replyTo int, caption, parseMode string) *VideoSender {
	return &VideoSender{bot: bot, chatIdInt: 0, chatidString: chatId, replyTo: replyTo, caption: caption, parseMode: parseMode}
}

/*Returns an AudioSender which has several methods for sending a audio. This method is only used for sending a audio to all types of chat except channels. To send a audio to a channel use "SendAudioToChannel" method.
To ignore int arguments pass 0 and to ignore string arguments pass empty string ("")

---------------------------------

Official telegram doc :

Use this method to send audio files, if you want Telegram clients to display them in the music player. Your audio must be in the .MP3 or .M4A format. On success, the sent Message is returned. Bots can currently send audio files of up to 50 MB in size, this limit may be changed in the future.

For sending voice messages, use the sendVoice method instead.*/
func (bot *Bot) SendAudio(chatId, replyTo int, caption, parseMode string) *AudioSender {
	return &AudioSender{bot: bot, chatIdInt: chatId, chatidString: "", replyTo: replyTo, caption: caption, parseMode: parseMode}
}

/*Returns a AudioSender which has several methods for sending a audio. This method is only used for sending a audio to a channels.
To ignore int arguments pass 0 and to ignore string arguments pass empty string ("")

---------------------------------

Official telegram doc :

Use this method to send audio files, if you want Telegram clients to display them in the music player. Your audio must be in the .MP3 or .M4A format. On success, the sent Message is returned. Bots can currently send audio files of up to 50 MB in size, this limit may be changed in the future.

For sending voice messages, use the sendVoice method instead.*/
func (bot *Bot) SendAudioToChannel(chatId string, replyTo int, caption, parseMode string) *AudioSender {
	return &AudioSender{bot: bot, chatIdInt: 0, chatidString: chatId, replyTo: replyTo, caption: caption, parseMode: parseMode}
}

/*Returns a DocumentSender which has several methods for sending a document. This method is only used for sending a document to all types of chat except channels. To send a audio to a channel use "SendDocumentToChannel" method.
To ignore int arguments pass 0 and to ignore string arguments pass empty string ("")

---------------------------------

Official telegram doc :

Use this method to send general files. On success, the sent Message is returned. Bots can currently send files of any type of up to 50 MB in size, this limit may be changed in the future.*/
func (bot *Bot) SendDocument(chatId, replyTo int, caption, parseMode string) *DocumentSender {
	return &DocumentSender{bot: bot, chatIdInt: chatId, chatidString: "", replyTo: replyTo, caption: caption, parseMode: parseMode}
}

/*Returns an AnimationSender

/*Returns a DocumentSender which has several methods for sending a document. This method is only used for sending a document to a channels.
To ignore int arguments pass 0 and to ignore string arguments pass empty string ("")

---------------------------------

Official telegram doc :

Use this method to send general files. On success, the sent Message is returned. Bots can currently send files of any type of up to 50 MB in size, this limit may be changed in the future.*/
func (bot *Bot) SendDocumentToChannel(chatId string, replyTo int, caption, parseMode string) *DocumentSender {
	return &DocumentSender{bot: bot, chatIdInt: 0, chatidString: chatId, replyTo: replyTo, caption: caption, parseMode: parseMode}
}

/*Returns an AnimationSender which has several methods for sending an animation. This method is only used for sending an animation to all types of chat except channels. To send a audio to a channel use "SendAnimationToChannel" method.
To ignore int arguments pass 0 and to ignore string arguments pass empty string ("")

---------------------------------

Official telegram doc :

Use this method to send animation files (GIF or H.264/MPEG-4 AVC video without sound). On success, the sent Message is returned. Bots can currently send animation files of up to 50 MB in size, this limit may be changed in the future.*/
func (bot *Bot) SendAnimation(chatId int, replyTo int, caption, parseMode string) *AnimationSender {
	return &AnimationSender{chatIdInt: chatId, chatidString: "", replyTo: replyTo, bot: bot, caption: caption, parseMode: parseMode}
}

/*Returns an AnimationSender which has several methods for sending an animation. This method is only used for sending an animation to channels
To ignore int arguments pass 0 and to ignore string arguments pass empty string ("")

---------------------------------

Official telegram doc :

Use this method to send animation files (GIF or H.264/MPEG-4 AVC video without sound). On success, the sent Message is returned. Bots can currently send animation files of up to 50 MB in size, this limit may be changed in the future.*/
func (bot *Bot) SendAnimationToChannel(chatId string, replyTo int, caption, parseMode string) *AnimationSender {
	return &AnimationSender{chatIdInt: 0, chatidString: chatId, replyTo: replyTo, bot: bot, caption: caption, parseMode: parseMode}
}

/*Returns a VocieSender which has several methods for sending a voice. This method is only used for sending a voice to all types of chat except channels. To send a voice to a channel use "SendVoiceToChannel" method.
To ignore int arguments pass 0 and to ignore string arguments pass empty string ("")

---------------------------------

Official telegram doc :

Use this method to send audio files, if you want Telegram clients to display the file as a playable voice message. For this to work, your audio must be in an .OGG file encoded with OPUS (other formats may be sent as Audio or Document). On success, the sent Message is returned. Bots can currently send voice messages of up to 50 MB in size, this limit may be changed in the future.*/
func (bot *Bot) SendVoice(chatId int, replyTo int, caption, parseMode string) *VoiceSender {
	return &VoiceSender{chatIdInt: chatId, chatidString: "", replyTo: replyTo, bot: bot, caption: caption, parseMode: parseMode}
}

/*Returns an VoiceSender which has several methods for sending a voice. This method is only used for sending a voice to channels.
To ignore int arguments pass 0 and to ignore string arguments pass empty string ("")

---------------------------------

Official telegram doc :

Use this method to send audio files, if you want Telegram clients to display the file as a playable voice message. For this to work, your audio must be in an .OGG file encoded with OPUS (other formats may be sent as Audio or Document). On success, the sent Message is returned. Bots can currently send voice messages of up to 50 MB in size, this limit may be changed in the future.*/
func (bot *Bot) SendVoiceToChannel(chatId string, replyTo int, caption, parseMode string) *VoiceSender {
	return &VoiceSender{chatIdInt: 0, chatidString: chatId, replyTo: replyTo, bot: bot, caption: caption, parseMode: parseMode}
}

/*Returns a VideoNoteSender which has several methods for sending a video note. This method is only used for sending a video note to all types of chat except channels. To send a video note to a channel use "SendVideoNoteToChannel" method.
To ignore int arguments pass 0 and to ignore string arguments pass empty string ("")

---------------------------------

Official telegram doc :

As of v.4.0, Telegram clients support rounded square mp4 videos of up to 1 minute long. Use this method to send video messages. On success, the sent Message is returned.*/
func (bot *Bot) SendVideoNote(chatId int, replyTo int, caption, parseMode string) *VideoNoteSender {
	return &VideoNoteSender{chatIdInt: chatId, chatidString: "", replyTo: replyTo, bot: bot, caption: caption, parseMode: parseMode}
}

/*Returns an VideoNoteSender which has several methods for sending a video note. This method is only used for sending a video note to channels.
To ignore int arguments pass 0 and to ignore string arguments pass empty string ("")

---------------------------------

Official telegram doc :

As of v.4.0, Telegram clients support rounded square mp4 videos of up to 1 minute long. Use this method to send video messages. On success, the sent Message is returned.*/
func (bot *Bot) SendVideoNoteToChannel(chatId string, replyTo int, caption, parseMode string) *VideoNoteSender {
	return &VideoNoteSender{chatIdInt: 0, chatidString: chatId, replyTo: replyTo, bot: bot, caption: caption, parseMode: parseMode}
}

/*To ignore replyTo argument, pass 0.*/
func (bot *Bot) CreateAlbum(replyTo int) *MediaGroup {
	return &MediaGroup{replyTo: replyTo, bot: bot, media: make([]objs.InputMedia, 0), files: make([]*os.File, 0)}
}

/*Sends a venue to all types of chat but channels. To send it to channels use "SendVenueToChannel" method.

---------------------------------

Official telegram doc :

Use this method to send information about a venue. On success, the sent Message is returned.*/
func (bot *Bot) SendVenue(chatId, replyTo int, latitude, longitude float32, title, address string, silent bool) (*objs.SendMethodsResult, error) {
	return bot.apiInterface.SendVenue(
		chatId, "", latitude, longitude, title, address, "", "", "", "", replyTo, silent, false, nil,
	)
}

/*Sends a venue to a channel.

---------------------------------

Official telegram doc :

Use this method to send information about a venue. On success, the sent Message is returned.*/
func (bot *Bot) SendVenueTOChannel(chatId string, replyTo int, latitude, longitude float32, title, address string, silent bool) (*objs.SendMethodsResult, error) {
	return bot.apiInterface.SendVenue(
		0, chatId, latitude, longitude, title, address, "", "", "", "", replyTo, silent, false, nil,
	)
}

/*Sends a contact to all types of chat but channels. To send it to channels use "SendContactToChannel" method.

---------------------------------

Official telegram doc :

Use this method to send phone contacts. On success, the sent Message is returned.*/
func (bot *Bot) SendContact(chatId, replyTo int, phoneNumber, firstName, lastName string, silent bool) (*objs.SendMethodsResult, error) {
	return bot.apiInterface.SendContact(
		chatId, "", phoneNumber, firstName, lastName, "", replyTo, silent, false, nil,
	)
}

/*Sends a contact to a channel.

---------------------------------

Official telegram doc :

Use this method to send phone contacts. On success, the sent Message is returned.*/
func (bot *Bot) SendContactToChannel(chatId string, replyTo int, phoneNumber, firstName, lastName string, silent bool) (*objs.SendMethodsResult, error) {
	return bot.apiInterface.SendContact(
		0, chatId, phoneNumber, firstName, lastName, "", replyTo, silent, false, nil,
	)
}

/*Creates a poll for all types of chat but channels. To create a poll for channels use "CreatePollForChannel" method.

The poll type can be "regular" or "quiz"*/
func (bot *Bot) CreatePoll(chatId int, question, pollType string) (*Poll, error) {
	if pollType != "quiz" && pollType != "regular" {
		return nil, errors.New("poll type invalid : " + pollType)
	}
	return &Poll{bot: bot, pollType: pollType, chatIdInt: chatId, question: question, options: make([]string, 0)}, nil
}

/*Creates a poll for a channel.

The poll type can be "regular" or "quiz"*/
func (bot *Bot) CreatePollForChannel(chatId, question, pollType string) (*Poll, error) {
	if pollType != "quiz" && pollType != "regular" {
		return nil, errors.New("poll type invalid : " + pollType)
	}
	return &Poll{bot: bot, pollType: pollType, chatIdString: chatId, question: question, options: make([]string, 0)}, nil
}

/*Sends a dice message to all types of chat but channels. To send it to channels use "SendDiceToChannel" method.

Available emojies : “🎲”, “🎯”, “🏀”, “⚽”, “🎳”, or “🎰”.

---------------------------------

Official telegram doc :

Use this method to send an animated emoji that will display a random value. On success, the sent Message is returned*/
func (bot *Bot) SendDice(chatId, replyTo int, emoji string, silent bool) (*objs.SendMethodsResult, error) {
	return bot.apiInterface.SendDice(
		chatId, "", emoji, replyTo, silent, false, nil,
	)
}

/*Sends a dice message to a channel.

Available emojies : “🎲”, “🎯”, “🏀”, “⚽”, “🎳”, or “🎰”.

---------------------------------

Official telegram doc :

Use this method to send an animated emoji that will display a random value. On success, the sent Message is returned*/
func (bot *Bot) SendDiceToChannel(chatId string, replyTo int, emoji string, silent bool) (*objs.SendMethodsResult, error) {
	return bot.apiInterface.SendDice(
		0, chatId, emoji, replyTo, silent, false, nil,
	)
}

/*Sends a chat action message to all types of chat but channels. To send it to channels use "SendChatActionToChannel" method.

---------------------------------

Official telegram doc :

Use this method when you need to tell the user that something is happening on the bot's side. The status is set for 5 seconds or less (when a message arrives from your bot, Telegram clients clear its typing status). Returns True on success.

Example: The ImageBot needs some time to process a request and upload the image. Instead of sending a text message along the lines of “Retrieving image, please wait…”, the bot may use sendChatAction with action = upload_photo. The user will see a “sending photo” status for the bot.

We only recommend using this method when a response from the bot will take a noticeable amount of time to arrive.

action is the type of action to broadcast. Choose one, depending on what the user is about to receive: typing for text messages, upload_photo for photos, record_video or upload_video for videos, record_voice or upload_voice for voice notes, upload_document for general files, choose_sticker for stickers, find_location for location data, record_video_note or upload_video_note for video notes.*/
func (bot *Bot) SendChatAction(chatId int, action string) (*objs.SendMethodsResult, error) {
	return bot.apiInterface.SendChatAction(chatId, "", action)
}

/*Sends a chat action message to a channel.

---------------------------------

Official telegram doc :

Use this method when you need to tell the user that something is happening on the bot's side. The status is set for 5 seconds or less (when a message arrives from your bot, Telegram clients clear its typing status). Returns True on success.

Example: The ImageBot needs some time to process a request and upload the image. Instead of sending a text message along the lines of “Retrieving image, please wait…”, the bot may use sendChatAction with action = upload_photo. The user will see a “sending photo” status for the bot.

We only recommend using this method when a response from the bot will take a noticeable amount of time to arrive.

action is the type of action to broadcast. Choose one, depending on what the user is about to receive: typing for text messages, upload_photo for photos, record_video or upload_video for videos, record_voice or upload_voice for voice notes, upload_document for general files, choose_sticker for stickers, find_location for location data, record_video_note or upload_video_note for video notes.*/
func (bot *Bot) SendChatActionToChannel(chatId, action string) (*objs.SendMethodsResult, error) {
	return bot.apiInterface.SendChatAction(0, chatId, action)
}

/*Sends a location (not live) to all types of chats but channels. To send it to channel use "SendLocationToChannel" method.

You can not use this methods to send a live location. To send a live location use AdvancedBot.

---------------------------------

Official telegram doc :

Use this method to send point on the map. On success, the sent Message is returned.*/
func (bot *Bot) SendLocation(chatId int, silent bool, latitude, longitude, accuracy float32, replyTo int) (*objs.SendMethodsResult, error) {
	return bot.apiInterface.SendLocation(
		chatId, "", latitude, longitude, accuracy, 0, 0, 0, replyTo, silent, false, nil,
	)
}

/*Sends a location (not live) to a channel.

You can not use this methods to send a live location. To send a live location use AdvancedBot.

---------------------------------

Official telegram doc :

Use this method to send point on the map. On success, the sent Message is returned.*/
func (bot *Bot) SendLocationToChannel(chatId string, silent bool, latitude, longitude, accuracy float32, replyTo int) (*objs.SendMethodsResult, error) {
	return bot.apiInterface.SendLocation(
		0, chatId, latitude, longitude, accuracy, 0, 0, 0, replyTo, silent, false, nil,
	)
}

/*Gets the given user profile photos.

"userId" argument is required. Other arguments are optinoal and to ignore them pass 0.

---------------------------------

Official telegram doc :

Use this method to get a list of profile pictures for a user. Returns a UserProfilePhotos object.*/
func (bot *Bot) GetUserProfilePhotos(userId, offset, limit int) (*objs.ProfilePhototsResult, error) {
	return bot.apiInterface.GetUserProfilePhotos(userId, offset, limit)
}

/*Gets a file from telegram server. If it is successful the File object is returned.

If "download option is true, the file will be saved into the given file and if the given file is nil file will be saved in the same name as it has been saved in telegram servers.*/
func (bot *Bot) GetFile(fileId string, download bool, file *os.File) (*objs.File, error) {
	res, err := bot.apiInterface.GetFile(fileId)
	if err != nil {
		return nil, err
	}
	if download {
		err2 := bot.apiInterface.DownloadFile(res.Result, file)
		if err2 != nil {
			return &res.Result, err2
		}
	}
	return &res.Result, nil
}

/*Creates and returnes a ChatManager for groups and other chats witch an integer id.

To manage supergroups and channels which have usernames use "GetChatManagerByUsername".*/
func (bot *Bot) GetChatManagerById(chatId int) *ChatManager {
	return &ChatManager{bot: bot, chatIdInt: chatId, chatIdString: ""}
}

/*Creates and returnes a ChatManager for supergroups and channels which have usernames

To manage groups and other chats witch an integer id use "GetChatManagerById".*/
func (bot *Bot) GetChatManagerByUsrename(chatId int) *ChatManager {
	return &ChatManager{bot: bot, chatIdInt: chatId, chatIdString: ""}
}

/*Use this method to send answers to callback queries sent from inline keyboards. The answer will be displayed to the user as a notification at the top of the chat screen or as an alert. On success, True is returned.

Alternatively, the user can be redirected to the specified Game URL. For this option to work, you must first create a game for your bot via @Botfather and accept the terms. Otherwise, you may use links like t.me/your_bot?start=XXXX that open your bot with a parameter.*/
func (bot *Bot) AnswerCallbackQuery(callbackQueryId, text string, showAlert bool) (*objs.LogicalResult, error) {
	return bot.apiInterface.AnswerCallbackQuery(callbackQueryId, text, "", showAlert, 0)
}

/*Returnes a command manager which has several method for manaing bot commands.*/
func (bot *Bot) GetCommandManager() *CommandsManager {
	return &CommandsManager{bot: bot}
}

/*Returnes a MessageEditor for a chat with id which has several methods for editing messages.

To edit messages in a channel or a chat with username, use "GetMsgEditorWithUN"*/
func (bot *Bot) GetMsgEditor(chatId int) *MessageEditor {
	return &MessageEditor{bot: bot, chatIdInt: chatId}
}

/*Returnes a MessageEditor for a chat with username which has several methods for editing messages.*/
func (bot *Bot) GetMsgEditorWithUN(chatId string) *MessageEditor {
	return &MessageEditor{bot: bot, chatIdInt: 0, chatIdString: chatId}
}

/*Stops the bot*/
func (bot *Bot) Stop() {
	bot.apiInterface.StopUpdateRoutine()
	*bot.pollRoutineChannel <- true
}

/*Returns and advanced version which gives more customized functions to iteract with the bot*/
func (bot *Bot) AdvancedMode() *AdvancedBot {
	return &AdvancedBot{Bot: bot}
}

func (bot *Bot) startPollUpdateRoutine() {
loop:
	for {
		select {
		case <-*bot.pollRoutineChannel:
			break loop
		default:
			poll := <-*bot.pollUpdateChannel
			pl := Polls[poll.Id]
			if pl == nil {
				logger.Logger.Println("Could not update poll `" + poll.Id + "`. Not found in the Polls map")
				continue
			}
			err3 := pl.Update(poll)
			if err3 != nil {
				logger.Logger.Println("Could not update poll `" + poll.Id + "`." + err3.Error())
			}
		}
	}
}

/*Return a new bot instance with the specified configs*/
func NewBot(cfg *cfg.BotConfigs) (*Bot, error) {
	api, err := tba.CreateInterface(cfg)
	if err != nil {
		return nil, err
	}
	ch := make(chan bool)
	return &Bot{botCfg: cfg, apiInterface: api, updateChannel: api.GetUpdateChannel(), pollUpdateChannel: api.GetPollUpdateChannel(), pollRoutineChannel: &ch}, nil
}
